# Social Media DB

A database system for a social media platform, designed for efficiency and scalability. Supports user accounts, posts, interactions, messaging, and social relationships.
I (lmProgramming/Mikołaj Kubś) am just one of the four contributors to this project for the "Database Design" course in PWr.

The idea of the course was to first design a database in multiple formal stages, then implement it using SQL. We used GORM and Go for this part. Next part of the course was to do the same thing again, but in NoSQL. We used MongoDB then. Refer to the [report](docs/main.pdf) for more info.

## Features

User Management: Registration, login, role-based access (admin, user, guest).
Content Interaction: Posts, likes, comments, hashtags, media storage via external links.
Messaging: Private conversations and group chats.
Friend System: Friend requests, followers, and relationship tracking.
Events & Pages: Organization and management of public pages and events.
Reporting: SQL queries for data insights (e.g., trending hashtags, most active users).
Index Optimization: Use of B-Tree, Hash, and BRIN indexes for performance.
Soft Delete: Retains deleted data for recovery.
ORM Support: Implemented using GORM for PostgreSQL.
NoSQL Integration: MongoDB-based alternative for scalability.
Infrastructure: Dockerized database deployment with automated scripts.

## Tech Stack

Database: PostgreSQL, MongoDB
ORM: GORM (Go)
Data Generation: Faker (Python & Go)
Backend Support: Makefile automation, Dockerized setup
