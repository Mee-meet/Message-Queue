# Message Consumer

## Overview

The message consumer code connects to a RabbitMQ broker, declares a queue, consumes encoded messages from the queue, decodes the messages, and processes the decoded message content. 

It can handle multiple message types like SMS, WhatsApp, email, etc by using the appropriate decoding library.

## Usage

To use the consumer:

1. Update the RabbitMQ connection string in the code if needed  
2. Build and run the consumer code
3. It will connect to RabbitMQ, declare a queue, and start consuming

## Functionality 

The main functions of the consumer code:

- Connect to RabbitMQ broker
- Open a channel 
- Declare a queue to consume from  
- Start consuming from the queue
- For each message consumed:
   - Decode message using appropriate library
   - Process decoded message content

## Configuration

The RabbitMQ connection string and queue name are currently hardcoded. These need to be updated to match your environment. 

The decoding library to use can be specified based on the message type.

## Dependencies

The consumer relies on:

- RabbitMQ client library
- Decoding libraries for each message type

## Logging 

Basic logging is included to track:

- Number of messages consumed 
- Message content
- Errors

Logging can be customized as needed.

## Next Steps

Some ways the consumer could be expanded:

- Parameterize configuration
- Persist decoded messages
- Add detailed metrics and monitoring  
- Support additional message types
- Integrate with external systems
