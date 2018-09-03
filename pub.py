from kafka import KafkaProducer
from kafka import KafkaConsumer
producer = KafkaProducer(bootstrap_servers='192.168.1.102:9092')
consumer = KafkaConsumer('response_conzumer',bootstrap_servers=['192.168.1.102:9092'])


producer.send('mob_number', value='0712263864')
for message in consumer:
	print(message.value)
	break
