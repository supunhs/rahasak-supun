from kafka import KafkaConsumer
from kafka import KafkaProducer

consumer = KafkaConsumer('mob_number',bootstrap_servers=['192.168.1.102:9092'])
producer = KafkaProducer(bootstrap_servers='192.168.1.102:9092')

for message in consumer:
	print ("value=%s" % (message.value))
	rule = re.compile(r'^(?:\+?44)?[07]\d{9,13}$')
	if rule.search(message.value):
		producer.send('response_conzumer', value = b'ok')
	else:
		producer.send('response_conzumer', value = b'wrong')
