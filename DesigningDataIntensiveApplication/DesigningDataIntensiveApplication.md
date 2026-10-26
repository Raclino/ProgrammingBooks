# Designing Data-Intensive Applications

**Authors:** Martin Kleppmann & Chris Riccomini  
**Edition:** 2nd Edition (2026)  
**Publisher:** O'Reilly Media

> Personal notes | Key concepts, architectural trade-offs, and system design fundamentals.

---

# Chapter 1 — Trade-Offs in Data Systems Architecture

## 1. Data-Intensive Applications

Modern applications rely on multiple data systems, each optimized for different workloads.

Common components:

- **Databases:** Store and retrieve persistent data (PostgreSQL).
- **Caches:** Speed up frequently accessed data (Redis).
- **Search indexes:** Enable efficient full-text search (Elasticsearch).
- **Message brokers / event streams:** Enable asynchronous communication (RabbitMQ, Kafka).
- **Analytical systems:** Process large datasets for reporting and analysis.

**Key idea:** No single technology is optimal for every workload. Architecture is about choosing the right trade-offs.

## 2. Operational vs. Analytical Systems

### OLTP — Online Transaction Processing

- Handles day-to-day application operations.
- Frequent, short reads and writes.
- Usually accesses a small number of records per query.
- Optimized for low latency, concurrency, and data integrity.
- Example: PostgreSQL storing users, orders, and payments.

### OLAP — Online Analytical Processing

- Processes large datasets for analytics and business intelligence.
- Complex queries scanning and aggregating many records.
- Often optimized for read-heavy workloads and columnar storage.
- Example: calculating monthly revenue across millions of transactions.

**Why separate them?**

Heavy analytical queries can consume CPU, memory, and disk I/O, degrading production database performance.

### Data Warehouses

Centralize data from multiple operational systems for analytical queries.

Typical pipeline:

`OLTP databases → ETL/ELT → Data Warehouse → BI / Analytics`

- **ETL:** Extract, Transform, Load.
- **ELT:** Extract, Load, Transform.
- **Data warehouse:** Structured, analytics-oriented data storage.
- **Data lake:** Stores large amounts of raw or semi-structured data, often using object storage.

## 3. Systems of Record vs. Derived Data

### System of Record (Source of Truth)

The authoritative version of data.

- Must prioritize durability and correctness.
- Losing it may cause irreversible data loss.
- Example: PostgreSQL storing customer orders.

### Derived Data

Data computed, copied, indexed, or transformed from other data.

Examples:

- Redis cache.
- Elasticsearch index.
- Precomputed recommendations.
- Materialized views.
- Analytical aggregates.

**Key property:** Derived data can often be rebuilt from the original data, provided the necessary source data and transformations remain available.

### Main Trade-Offs

- Derived data improves performance and supports specialized workloads.
- Maintaining multiple representations increases complexity.
- Synchronization delays may cause stale data.
- Cache invalidation and consistency become important.

**Important:** Rebuildable does not mean disposable without consequences. Losing a cache can overload the primary database (_cache stampede_).

## 4. Cloud vs. Self-Hosting

### Cloud Advantages

- No physical hardware management.
- Rapid provisioning and elasticity.
- Managed services reduce operational overhead.
- Global infrastructure and geographic distribution.
- Usage-based pricing can benefit unpredictable workloads.

### Cloud Trade-Offs

- **Vendor lock-in:** Dependence on provider-specific services.
- **Cost unpredictability:** Scaling, storage, and data transfer fees.
- **Reduced control:** Less access to underlying infrastructure.
- **Provider dependency:** Outages, policy changes, service availability.
- **Shared responsibility:** The provider secures the cloud infrastructure; customers remain responsible for their applications, configurations, and data.

### Self-Hosting Advantages

- Greater infrastructure control.
- Potentially lower costs for large, stable, predictable workloads.
- Less dependence on a specific cloud provider.

### Self-Hosting Trade-Offs

- Hardware acquisition and maintenance.
- Capacity planning and slower scaling.
- Dedicated operational expertise.
- Responsibility for redundancy, security, and disaster recovery.

**Key idea:** Compare total cost of ownership (TCO), not just server prices.

## 5. Cloud-Native Architecture

Cloud-native means designing applications to exploit cloud infrastructure characteristics, rather than simply hosting traditional applications in the cloud.

Common principles:

- Automated provisioning and deployments.
- Elastic scaling.
- Replaceable / ephemeral compute instances.
- Separation of compute lifecycle from persistent data.
- Managed infrastructure and services where appropriate.
- Resilience to infrastructure failures.

**Lift-and-shift ≠ cloud-native.**

Moving an unchanged monolith from an on-premises server to EC2 does not automatically make it cloud-native.

**Cloud-native ≠ microservices.** A well-designed monolith can also be cloud-native.

## 6. Single-Node vs. Distributed Systems

### Why Distribute a System?

- **Scalability:** Exceed the CPU, RAM, or storage capacity of one machine.
- **Fault tolerance:** Continue operating despite individual machine failures.
- **Geographic distribution:** Reduce latency by placing systems closer to users.
- **Isolation:** Scale or operate workloads independently.

### Distributed Systems Trade-Offs

Distribution introduces new failure modes:

- Network latency and timeouts.
- Partial failures.
- Network partitions.
- Data consistency and synchronization.
- More complex debugging and observability.
- Coordination between machines.

**Critical concept:** A timeout does not prove that an operation failed.

A request may have succeeded even if its response was lost. Blind retries can therefore cause duplicate operations.

Useful mechanisms include idempotency keys, retries with backoff, and deduplication.

**Key idea:** Prefer a single-node system when it meets the requirements. Distribution should solve a real problem, not hypothetical future scale.

## 7. Monoliths vs. Microservices

### Monolith

- Single deployable application.
- Simpler development, testing, debugging, and deployment.
- Easier local transactions and data consistency.
- Can be organized as a **modular monolith** with clear domain boundaries.

### Microservices

- Independently deployable services.
- Allow independent scaling and technology choices.
- Support team autonomy and domain ownership.
- Can isolate failures and resource consumption.

### Microservices Trade-Offs

- Network communication replaces local function calls.
- Distributed transactions become difficult.
- Partial failures require explicit handling.
- More complex CI/CD, monitoring, tracing, and deployment.
- Cross-service data consistency becomes challenging.

Example:

`Order Service → Payment Service`

An order may be created while its payment fails. The architecture must handle this intermediate state.

**Key idea:** Start with a modular monolith unless independent deployment, scaling, or organizational requirements justify distribution.

## 8. Serverless

Serverless abstracts infrastructure management away from application developers.

Examples: AWS Lambda, Azure Functions.

### Advantages

- Automatic scaling.
- Reduced infrastructure management.
- Usage-based billing.
- Well suited to event-driven and intermittent workloads.

### Limitations

- Cold starts.
- Execution time and resource limits.
- Provider-specific integrations and vendor lock-in.
- Potentially higher costs for sustained workloads.
- More complex debugging and observability.

**Serverless ≠ no servers.** Servers still execute the code; the cloud provider manages their lifecycle.

## 9. Cloud Computing vs. Supercomputing

- **Cloud computing:** General-purpose, elastic infrastructure supporting many independent workloads.
- **Supercomputing / HPC:** Tightly coordinated computation optimized for high-performance parallel processing.
- Cloud workloads often tolerate variable latency and partial failures; tightly coupled HPC workloads may require low-latency, high-bandwidth interconnects.

**Key idea:** Different computational workloads require different infrastructure trade-offs.

## 10. Data Systems, Law, and Society

Architecture decisions also have legal and societal consequences.

- **Privacy:** Limit collection and exposure of personal data.
- **Data residency:** Regulations may constrain where data is stored or processed.
- **Security:** Protect confidentiality and integrity.
- **Governance:** Define ownership, retention, access, and deletion policies.
- **Societal impact:** Consider how data collection and automated decisions affect people.

**Key idea:** Technical efficiency is not the only architectural requirement.

---

## Chapter 1 — Key Takeaways

1. **There is no universally best architecture.** Every design involves trade-offs.
2. **OLTP and OLAP serve different workloads** and often benefit from separate infrastructure.
3. **Systems of record are authoritative; derived data is often reconstructible.**
4. **Cloud simplifies infrastructure management but introduces cost, control, and dependency trade-offs.**
5. **Cloud-native is an architectural approach, not simply a hosting location.**
6. **Distributed systems improve scalability and resilience but introduce partial failures and coordination complexity.**
7. **Microservices should solve real scaling or organizational problems, not be the default.**
8. **Choose the simplest architecture that satisfies current and reasonably expected requirements.**
