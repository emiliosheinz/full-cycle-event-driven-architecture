import kafka from 'kafka-node';
import { Message } from './Message';
import { Topic } from './Topic';

type MessageValue = { name: Message; payload: unknown };

export class KafkaConsumer {
  private client: kafka.KafkaClient;

  constructor() {
    this.client = new kafka.KafkaClient({ kafkaHost: 'kafka:29092' });
  }

  public async toBeReady(topic: Topic) {
    return new Promise<void>(resolve => {
      const interval = setInterval(() => {
        this.client.topicExists([topic], error => {
          if (!error) {
            clearInterval(interval);
            resolve();
          } else {
            console.log(`Waiting for ${topic} topic to be ready...`);
          }
        });
      }, 3_000);
    });
  }

  public consume(topic: Topic, onMessage: (value: MessageValue) => void) {
    const consumer = new kafka.Consumer(this.client, [{ topic: topic }], {
      autoCommit: false,
    });
    consumer.on('message', message => {
      const { Name: name, Payload: payload } = JSON.parse(
        message.value as string,
      );
      if (!name || !payload) {
        throw new Error('Unporcessable message');
      }
      onMessage({ name, payload });
    });
  }
}
