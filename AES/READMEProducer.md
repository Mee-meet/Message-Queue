# Producer Code README

## Overview

The code does the following:

- Connects to a RabbitMQ server
- Opens a channel
- Declares a queue named "SMSQUEUE" 
- Encodes a sample message using the dongle library
- Publishes the message to the "SMSQUEUE" queue

## Usage

To use this producer:

  - Update the RabbitMQ connection string if needed
  - Build and run the code  
  - The encoded sample message will be published to the "SMSQUEUE" queue

The sample message encodes a string "SM2;+91912287226272" using Base64 encoding. This can be updated to encode and send any message content needed. 

The queue, exchange, and routing key are hardcoded for simplicity but can be parameterized if needed.

## Dependencies

- RabbitMQ server
- amqp091-go RabbitMQ client library
- dongle encoding library

## Configuration

The RabbitMQ connection string and queue name are currently hardcoded and will need to be updated to match your environment.

## Logging

Basic logging is included to print out message publication events and errors.

Let me know if you would like me to explain or expand on any part of this README!
