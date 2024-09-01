import kafka from 'kafka-node';
import { Message } from './Message';

type MessageValue = { name: Message; payload: unknown };

export class KafkaConsumer {
  private client: kafka.KafkaClient;

  constructor() {
    this.client = new kafka.KafkaClient({ kafkaHost: 'kafka:29092' });
  }

  public consume(topic: string, onMessage: (value: MessageValue) => void) {
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
