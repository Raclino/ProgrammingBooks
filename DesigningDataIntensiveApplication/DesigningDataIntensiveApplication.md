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

## Chapter 1 — Key Takeaways

1. **There is no universally best architecture.** Every design involves trade-offs.
2. **OLTP and OLAP serve different workloads** and often benefit from separate infrastructure.
3. **Systems of record are authoritative; derived data is often reconstructible.**
4. **Cloud simplifies infrastructure management but introduces cost, control, and dependency trade-offs.**
5. **Cloud-native is an architectural approach, not simply a hosting location.**
6. **Distributed systems improve scalability and resilience but introduce partial failures and coordination complexity.**
7. **Microservices should solve real scaling or organizational problems, not be the default.**
8. **Choose the simplest architecture that satisfies current and reasonably expected requirements.**

---

# Chapter 2 — Defining Nonfunctional Requirements

## 1. Functional vs. Nonfunctional Requirements

**Functional requirements:** What the system must do (features, business logic, APIs).

**Nonfunctional requirements (NFRs):** How well the system must operate.

Main categories:

- **Reliability:** Continue working correctly despite faults.
- **Scalability:** Handle increasing load without unacceptable performance degradation.
- **Maintainability:** Make the system easy to operate, understand, and evolve.

Other NFRs include security, privacy, availability, performance, and compliance.

**Key idea:** Architecture is largely driven by nonfunctional requirements and their trade-offs.

## 2. Reliability

A reliable system continues providing its intended service correctly, even when things go wrong.

### Fault vs. Failure

- **Fault:** A component deviates from its expected behavior.
- **Failure:** The overall system stops delivering the required service.

A fault does not necessarily cause a failure if the system can tolerate it.

Example: A Kubernetes pod crashes (fault), but another replica continues serving requests (no user-visible failure).

### Three Types of Faults

**Hardware faults**

- Disk failures, machine crashes, power outages.
- Mitigations: replication, redundancy, backups, failover.

**Software faults**

- Bugs, memory leaks, resource exhaustion, unexpected interactions.
- Mitigations: testing, isolation, gradual deployments, monitoring, rollback.

**Human errors**

- Misconfigurations, accidental deletion, incorrect deployments.
- Mitigations: automation, access controls, reviews, safe defaults, recovery procedures.

### Fault Tolerance

Design systems assuming individual components will eventually fail.

Common techniques:

- Redundancy and replication.
- Automatic failover.
- Retries with exponential backoff.
- Timeouts and circuit breakers.
- Graceful degradation.
- Backups and disaster recovery.

**Important:** Replication improves availability but does not replace backups. Accidental deletion or corrupted data may be replicated too.

## 3. Availability and Reliability Metrics

**Availability:** Proportion of time a system is operational and accessible.

Typical availability targets:

| Availability | Approx. maximum downtime/year |
| ------------ | ----------------------------- |
| 99%          | 3.65 days                     |
| 99.9%        | 8.76 hours                    |
| 99.99%       | 52.6 minutes                  |
| 99.999%      | 5.26 minutes                  |

Availability alone does not guarantee correctness or acceptable latency.

### SLI, SLO, SLA

- **SLI (Service Level Indicator):** Measured service behavior (latency, availability, error rate).
- **SLO (Service Level Objective):** Internal target for an SLI.
- **SLA (Service Level Agreement):** Contractual commitment, potentially with penalties.

Example:

- SLI: Percentage of HTTP requests completing successfully.
- SLO: 99.9% successful requests over 30 days.
- SLA: Contractual availability commitment of 99.5%.

**Error budget:** The amount of unreliability permitted by an SLO.

Example: A 99.9% success-rate SLO allows 0.1% unsuccessful requests within the measurement window.

## 4. Scalability

Scalability describes how well a system handles increasing load.

A system is not simply "scalable" or "not scalable": its scalability depends on its workload and architecture.

### Load Parameters

Metrics describing the workload imposed on a system.

Examples:

- Requests per second (RPS).
- Concurrent active users.
- Database reads/writes per second.
- Messages processed per second.
- Dataset size.
- Cache hit ratio.
- Fan-out per event.

**Key idea:** Define the load before evaluating scalability.

### Performance Metrics

**Latency:** Time a request spends waiting for a response.

**Response time:** Total time observed by the client, including processing and network delays.

**Throughput:** Amount of work completed per unit of time (requests/sec, MB/sec, events/sec).

### Percentiles

Averages can hide poor performance experienced by a minority of users.

- **p50:** Median response time.
- **p95:** 95% of requests complete within this time.
- **p99:** 99% of requests complete within this time.
- **p99.9:** Captures extreme tail latency.

Example:

`p50 = 80ms | p95 = 250ms | p99 = 2.5s`

Most requests are fast, but the slowest 1% take at least 2.5 seconds.

**Tail latency matters:** In distributed systems, one slow dependency can delay an entire request.

### Scaling Strategies

**Vertical scaling (scale up)**

- Add CPU, RAM, or storage to an existing machine.
- Simple, but limited by hardware and cost.

**Horizontal scaling (scale out)**

- Add machines and distribute the workload.
- Better potential capacity and fault tolerance.
- Introduces network, coordination, and consistency complexity.

**Elasticity:** Automatically adjusting resources to changing demand.

### Fan-Out: Social Feed Example

**Fan-out on write (push):**

- Precompute followers' feeds when a user posts.
- Fast reads, expensive writes for popular accounts.

**Fan-out on read (pull):**

- Assemble feeds when followers request them.
- Cheaper writes, potentially expensive reads.

**Hybrid:** Push for ordinary accounts; pull for extremely popular accounts.

**Key idea:** Workload characteristics determine the appropriate architecture.

## 5. Maintainability

Maintainability is the ability to operate, understand, and modify a system efficiently over time.

### Operability

Make it easy for operators to keep the system running.

- Monitoring, logging, tracing, and alerting.
- Automated deployments and recovery.
- Clear documentation and runbooks.
- Capacity planning and incident management.

### Simplicity

Reduce unnecessary and accidental complexity.

- Clear abstractions and module boundaries.
- Minimize hidden dependencies.
- Avoid unnecessary distributed components.
- Prefer understandable designs over clever implementations.

**Accidental complexity:** Complexity introduced by implementation choices rather than the inherent business problem.

**Key idea:** Simplicity does not necessarily mean fewer lines of code.

### Evolvability

Make the system easy to adapt to changing requirements.

- Modular architecture.
- Well-defined interfaces.
- Automated tests.
- Backward-compatible API/schema changes.
- Loose coupling and clear ownership.

**Key idea:** Requirements change; architecture should support change without excessive cost or risk.

## 6. Architectural Trade-Offs

Nonfunctional requirements often conflict.

Examples:

- Stronger consistency may increase latency or reduce availability during certain failures.
- Replication improves fault tolerance but increases operational complexity.
- Caching improves read performance but introduces invalidation and staleness.
- Microservices enable independent scaling but complicate debugging and transactions.
- More redundancy increases availability but also infrastructure cost.

**Key principle:** Optimize for explicit requirements, not abstract ideals such as "maximum scalability."

## Chapter 2 — Key Takeaways

1. **Reliability:** Tolerate faults without causing unacceptable system failures.
2. **Scalability:** Understand the workload before choosing a scaling strategy.
3. **Maintainability:** Optimize for operability, simplicity, and evolvability.
4. **Fault ≠ Failure:** Component failures should not necessarily become user-visible failures.
5. **SLI → SLO → SLA:** Measure behavior, set objectives, define commitments.
6. **p95/p99 matter:** Average latency alone hides tail performance problems.
7. **Scale up vs. scale out:** Simplicity versus distributed capacity and complexity.
8. **Architecture is about trade-offs:** Define measurable requirements before choosing technologies.

---
# Chapter 3 — Data Models and Query Languages

## 1. Data Models

Data models determine how information is represented, stored, queried, and related.

Main models:

- **Relational:** Tables, rows, foreign keys, joins.
- **Document:** JSON-like documents with nested structures.
- **Graph:** Nodes and edges representing relationships.

**Key idea:** The best model depends on data relationships and access patterns.

## 2. Relational Model

Data is organized into tables (relations), containing rows and columns.

- **Primary key (PK):** Uniquely identifies a row.
- **Foreign key (FK):** References another table.
- **Joins:** Combine related data across tables.
- **Normalization:** Reduces redundancy and update anomalies.

### Relationships

- **One-to-one (1:1):** User → Profile.
- **One-to-many (1:N):** User → Orders.
- **Many-to-many (N:M):** Playlists ↔ Songs, using a junction table.

### Normalization vs. Denormalization

**Normalization**
- Avoids unnecessary duplication.
- Improves consistency and simplifies updates.
- May require joins when reading related data.

**Denormalization**
- Duplicates or precomputes data to optimize reads.
- Reduces some joins.
- Increases storage and update complexity.
- Can introduce stale or inconsistent copies.

**Key idea:** Normalize by default; denormalize when access patterns justify it.

## 3. Document Model

Stores related information in self-contained documents, typically JSON-like structures.

Example:

```json
{
  "orderId": 123,
  "customerId": 42,
  "shippingAddress": {
    "city": "Paris"
  },
  "items": [
    { "productId": 1, "quantity": 2 }
  ]
}
```

### Advantages

- **Data locality:** Related information can be retrieved together.
- Natural representation of hierarchical data.
- Flexible document structures.
- Fewer joins for self-contained aggregates.

### Limitations

- Duplication can cause update anomalies.
- Many-to-many relationships are less natural.
- Cross-document queries and joins can be more complex.
- Large or frequently modified documents may be inefficient.

### When to Use

**Document model:** Data is naturally hierarchical and usually accessed together.

**Relational model:** Data contains many interconnected entities and requires flexible joins.

**Important:** Relational databases can also support JSON; relational and document features are not mutually exclusive.

## 4. Schema-on-Write vs. Schema-on-Read

### Schema-on-Write

Data is validated against a schema when written.

- Common in relational databases.
- Stronger structural guarantees.
- Schema evolution often involves migrations.

### Schema-on-Read

Structure is interpreted when data is read.

- Common in flexible document-oriented systems.
- Different document versions may coexist.
- Application code must handle missing or changed fields.

**Trade-off:** Early validation and consistency vs. flexibility and easier structural evolution.

## 5. Query Languages

### Declarative Queries

Specify **what** data is needed, not the exact execution steps.

Example:

```sql
SELECT name
FROM users
WHERE country = 'FR';
```

Advantages:

- Database optimizer chooses execution strategy.
- Easier to read and maintain.
- Query execution can improve without changing application code.
- Enables optimizations such as parallel execution.

### Imperative Queries

Specify **how** to retrieve or transform data step by step.

**Key idea:** Declarative languages separate application intent from execution strategy.

## 6. MapReduce

Programming model for processing large datasets.

- **Map:** Transform input records into intermediate key-value pairs.
- **Reduce:** Aggregate values associated with the same key.

Example: Counting occurrences of words across documents.

**Trade-offs:**
- Useful for distributed batch processing.
- Can involve significant intermediate data movement.
- Often less expressive and convenient than declarative query languages.

## 7. Graph Data Models

Graph models represent highly interconnected data.

### Property Graphs

- **Vertices (nodes):** Entities.
- **Edges:** Relationships between entities.
- Nodes and edges can have properties.
- Edges can have labels and directions.

Example:

`(User) ──FOLLOWS──> (User)`

Useful for social networks, recommendations, fraud detection, and knowledge graphs.

### Triple Stores

Represent information as:

`(Subject, Predicate, Object)`

Example:

`(Alice, livesIn, Paris)`

Often associated with RDF and semantic web technologies.

### Graph Query Languages

- **Cypher:** Pattern-matching queries for property graphs.
- **SPARQL:** Query language for RDF/triple stores.
- **Datalog:** Rule-based, declarative logic queries.

Graph queries express relationships and traversal patterns naturally.

**Key idea:** Graph models are particularly useful when relationships are central and traversal patterns are complex or evolving.

## 8. Choosing a Data Model

| Requirement | Common choice |
|---|---|
| Structured entities, complex joins | Relational |
| Hierarchical, self-contained data | Document |
| Highly interconnected entities | Graph |
| Flexible fields and nested objects | Document / JSON |
| Strong relational integrity | Relational |
| Relationship traversal | Graph |

These are tendencies, not strict limitations. Many modern databases support features from multiple models.


## Chapter 3 — Key Takeaways

1. **Choose data models based on relationships and access patterns**, not popularity.
2. **Relational models** handle interconnected data and joins naturally.
3. **Document models** favor hierarchical data and locality.
4. **Normalization** reduces redundancy; denormalization optimizes selected reads at the cost of update complexity.
5. **Schema-on-write** validates structure early; **schema-on-read** offers greater flexibility.
6. **Declarative queries** let the database optimize execution.
7. **MapReduce** separates distributed transformations into map and reduce stages.
8. **Graph models** represent complex relationships through nodes and edges.
9. **No model is universally superior:** each optimizes different workloads and introduces different trade-offs.

---
