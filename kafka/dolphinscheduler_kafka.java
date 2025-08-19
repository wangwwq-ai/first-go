// 自定义任务插件
public class KafkaTask implements Task {
    private KafkaConsumer<String, String> consumer;

    public void init() {
        Properties props = new Properties();
        props.put("bootstrap.servers", "localhost:9092");
        props.put("group.id", "test");
        props.put("enable.auto.commit", "false");
        props.put("key.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        props.put("value.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        consumer = new KafkaConsumer<>(props);
        consumer.subscribe(Arrays.asList("user_behavior"));
    }

    public void execute() {
        ConsumerRecords<String, String> records = consumer.poll(Duration.ofMillis(100));
        for (ConsumerRecord<String, String> record : records) {
            // 处理数据
            System.out.println(record.value());
        }
    }

    public void destroy() {
        consumer.close();
    }
}
