-- 创建Kafka引擎表
CREATE TABLE user_behavior_kafka (
    user_id UInt32,
    item_id UInt32,
    event_time DateTime
) ENGINE = Kafka('localhost:9092', 'user_behavior', 'group1', 'JSONEachRow');

-- 创建MergeTree表
CREATE TABLE user_behavior (
    user_id UInt32,
    item_id UInt32,
    event_time DateTime
) ENGINE = MergeTree()
ORDER BY (user_id, item_id, event_time);

-- 创建物化视图
CREATE MATERIALIZED VIEW user_behavior_mv TO user_behavior AS
SELECT user_id, item_id, event_time FROM user_behavior_kafka;
