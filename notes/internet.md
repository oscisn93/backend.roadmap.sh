# How Does the Internet Work #

These notes are based on the article from cs.fyi titled ["How does the Internet Work"](https://cs.fyi/guide/how-does-internet-work#how-the-internet-works-an-overview).
It is meant to teach the fundamentals of the internet for the purpose of providing domain specific knowledge to software developers who work with/on distributed digital systems.

## Background ##

The internet is a network of networks that allows and facilitates digital communication among systems across large ranges of distance. A network is thus the building block of the internet. A network is an interconnected group of computers and devices.

## Overview ##

The internet is not complicated, it is simple but complex. This is because it is built on a set of standardized protocols which defiine how information is exchanged between devices. The protocols alos ensure that data transmission is reliable and secure.These protocols are not complicated in and of themselves but there are various implementations and the interplay between not only different devices but systems and components of the internet create complex interplay that makes the internet complex as its scale grows. The core of the internet is a set of routers which are responsible for directing internet traffic. Internet traffic consists of small packets that are sent from a device to a router, and henceforth forwarded to as many additional routers as is necessary until they reach their final destination.

The most important protocols include the Internet Protocol (IP) which is responsible for routing packets to the correct destination, the Transmission Control Protocol (TCP) which is responsible for ensuring reliable, in-order delivery of packets. Additionally the User Datagram Protocol (UDP) serves as an alternative for when reliability of packet delivery is not as important.

## Glossary ##

- Packet: A small unit of data that is transmitted over the internet.
- Router: A device that directs packets of data between different networks.
- IP Address: A unique identifier assigned to each device on a network, used to route data to the correct destination.
- Domain Name: A human-readable name that is used to identify a website, such as google.com.
- DNS: The Domain Name System is responsible for translating domain names into IP addresses.
- HTTP: The Hypertext Transfer Protocol is used to transfer data between a client (such as a web browser) and a server (such as a website).
- HTTPS: An encrypted version of HTTP that is used to provide secure communication between a client and server.
- SSL/TLS: The Secure Sockets Layer and Transport Layer Security protocols are used to provide secure communication over the internet.

## Protocols &amp; The Internet ##

