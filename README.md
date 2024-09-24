# E-commerce Observable System

This project is an e-commerce backend management system developed in Golang, with observability capabilities configured for better monitoring and troubleshooting. The system is designed to provide a robust and scalable platform for managing e-commerce operations.

## Features

- Product management
- Order management
- Customer management
- Inventory management
- User roles and permissions
- Real-time monitoring and observability

## Tech Stack

- Backend: Golang
- Database: MySQL/PostgreSQL
- Containerization: Docker
- Orchestration: Kubernetes
- Observability: Prometheus, Grafana, ELK Stack (Elasticsearch, Logstash, Kibana)

## Project Structure
![15](https://github.com/user-attachments/assets/852ee728-d1fd-49a2-88aa-cce81822cfe9)

- `cart`: Shopping cart related functionality
- `cartapi`: Shopping cart API
- `category`: Product category management
- `common`: Common components and utilities
- `deploy-middleware`: Middleware deployment configurations
- `docker-compose`: Docker Compose configuration files
- `docker-elk`: ELK Stack Docker configurations
- `order`: Order management
- `payment`: Payment functionality
- `paymentapi`: Payment API
- `product`: Product management
- `user`: User management

## Requirements

- Golang 1.16+
- Docker
- Kubernetes
- Helm (optional)
- Prometheus
- Grafana
- ELK Stack

## Observability

This system integrates the following observability tools:
- Prometheus: For metrics collection and monitoring
- Grafana: For data visualization and dashboards
- ELK Stack: For log management and analysis

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgements

- All contributors who have helped with this project
- Open-source libraries and tools used in this project
