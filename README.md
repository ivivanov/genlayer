# Developers Challenge: Cyberpunk Data Security and Routing

## Objective
Evaluate your ability to solve complex algorithmic problems in a cyberpunk-themed setting, focusing on data security, allocation, and network routing. You will complete three coding exercises—Data Fragmentation and Reconstruction, Secure Data Fragment Allocation, and Optimal Network Routing—in a programming language specified by the recruiter or team member. Your solutions must be submitted via a Git repository within a 3-hour time limit, which begins when you receive the challenge. You may choose when to receive the challenge to accommodate your schedule.

## Instructions
1. **Language Requirement**: Solve all exercises in the programming language specified by the recruiter or team member (e.g., Python, Java, C++, etc.). If no language is specified, confirm with the team before starting.
2. **Time Limit**: You have **3 hours** to complete all three exercises, starting from the moment you receive this challenge. A timer will be enforced based on your submission timestamp.
3. **Submission**:
   - Create a **private Git repository** (e.g., on GitHub, GitLab, or Bitbucket) containing your solutions.
   - Structure the repository with separate directories for each exercise (e.g., `fragmentation`, `allocation`, `routing`).
   - Include a `README.md` in the root directory with:
     - Instructions to run each solution.
     - Any assumptions or dependencies.
     - A brief explanation of your approach for each exercise.
   - Grant access to the repository to the evaluation team (provide the team with the repository URL and necessary permissions).
   - Commit your code regularly to show progress, but ensure the final commit is made within the 3-hour window.
4. **Constraints**:
   - Do not use external libraries that obscure core logic (e.g., cryptographic libraries for hashing, graph libraries for routing).
   - Ensure solutions are efficient and handle large inputs or complex scenarios.
   - Test your code with provided examples and additional edge cases.
5. **Resources**: You may use standard documentation, IDEs, and debugging tools, but the work must be your own. Do not consult external individuals or AI tools during the challenge.

## Exercises

### Exercise 1: Data Fragmentation and Reconstruction
**Scenario**: In a cyberpunk world, data is fragmented and dispersed across storage units for security. Each fragment is stored with a hash value to verify integrity. You must reconstruct the original data while ensuring all fragments are valid.

**Task**: Implement a `reconstruct_data` function that reassembles data from fragments after verifying their integrity using a `simple_hash` function you design. The hash function must generate a fixed-length (30-character) hash without external cryptographic libraries.

**Input**:
- A dictionary `fragments` where:
  - Keys are integers (fragment sequence).
  - Values are dictionaries with `data` (string, the fragment) and `hash` (string, the fragment’s hash).
- The `simple_hash` function takes a string and returns a 30-character hash.

**Output**:
- A string of reconstructed data if all fragments are valid and in order.
- `"Error: Data integrity verification failed"` if any fragment’s hash does not match.

**Constraints**:
- `simple_hash` must minimize collisions and produce a 30-character output.
- Handle out-of-order fragments, missing fragments, or invalid hashes.

**Example**:
```python
fragments = {
    1: {'data': 'Hello', 'hash': simple_hash('Hello')},
    2: {'data': 'World', 'hash': simple_hash('World')},
    3: {'data': '!', 'hash': simple_hash('!')}
}
print(reconstruct_data(fragments))  # Output: "HelloWorld!"
```

**Evaluation**:
- Correctness: Accurate reconstruction and hash verification.
- Hash Design: Balances simplicity and collision resistance.
- Edge Cases: Handles missing, out-of-order, or tampered fragments.

### Exercise 2: Secure Data Fragment Allocation
**Scenario**: To thwart data breaches, sensitive information is split into fragments and distributed across data centers with varying risk factors. The goal is to minimize the maximum risk of any data center, where risk is the center’s base risk raised to the power of the number of fragments stored.

**Task**: Implement a `distributed_fragments` function that optimally distributes a given number of fragments across data centers to minimize the maximum risk.

**Input**:
- A list of integers `data_centers` (risk factors of each center).
- An integer `fragments` (total number of fragments to distribute).

**Output**:
- An integer representing the minimized maximum risk.

**Constraints**:
- Avoid external libraries for core logic.
- Handle large numbers of centers and fragments efficiently.

**Example**:
```python
data_centers = [10, 20, 30]
fragments = 5
print(distribute_fragments(data_centers, fragments))  # Output: 1000 (e.g., 10^3 = 1000 for 3 fragments in center 1)
```

**Evaluation**:
- Correctness: Accurately minimizes maximum risk.
- Efficiency: Scales with large inputs.
- Edge Cases: Handles high risk factors, few/many fragments, or single centers.

### Exercise 3: Optimal Network Routing
**Scenario**: In a cyberpunk city’s network, data must be routed efficiently through routers with varying latencies. Certain nodes can compress data, halving the latency of outgoing links, but this can be used only once per path. You must find the path with the minimum total latency.

**Task**: Implement a `find_minimum_latency_path` function that computes the shortest path from a source to a destination router, accounting for possible compression at specific nodes.

**Input**:
- `graph`: A dictionary where keys are router IDs and values are lists of tuples `(neighbor, latency)`.
- `compression_nodes`: A list of router IDs where compression (halving latency) can be applied.
- `source`: The source router ID.
- `destination`: The destination router ID.

**Output**:
- An integer representing the minimum total latency.

**Constraints**:
- Avoid external libraries for core logic.
- Handle complex graphs efficiently.

**Example**:
```python
graph = {
    'A': [('B', 10), ('C', 20)],
    'B': [('D', 15)],
    'C': [('D', 30)],
    'D': []
}
compression_nodes = ['B', 'C']
source = 'A'
target = 'D'
print(find_minimum_latency_path(graph, compression_nodes, source, target))  # Output: 17.5 (e.g., A->B->D with compression at B)
```

**Evaluation**:
- Correctness: Accurately computes minimum latency with compression.
- Efficiency: Handles heran Optimizes for complex graphs.
- Edge Cases: Handles disconnected paths, single nodes, or no compression nodes.

## Submission Guidelines
- **Repository Structure**:
  ```
  developers-challenge/
  ├── fragmentation/
  │   └── solution.<ext>  # e.g., solution.py, solution.java
  ├── allocation/
  │   └── solution.<ext>
  ├── routing/
  │   └── solution.<ext>
  └── README.md
  ```
- **README Requirements**:
  - Instructions to compile/run each solution (e.g., `python fragmentation/solution.py`).
  - Dependencies (e.g., language version, standard libraries).
  - Brief explanation of your approach for each exercise (1-2 sentences per exercise).
  - Any assumptions (e.g., input validation, edge case handling).
- **Commit History**: Regular commits to show progress (e.g., one per exercise or major feature).
- **Access**: Share the repository URL with the evaluation team, ensuring read access.

## Evaluation Criteria
Your submission will be evaluated based on:

### Correctness (30%)
- All three exercises produce correct outputs for provided examples and hidden test cases.
- Hash verification (Exercise 1), risk minimization (Exercise 2), and latency optimization (Exercise 3) are accurate.

### Algorithmic Efficiency (25%)
- Solutions scale efficiently for large inputs (e.g., many fragments, data centers, or graph nodes).
- Time and space complexity are reasonable (e.g., O(n log n) for routing, not O(n^3)).

### Code Quality (20%)
- Code is well-structured, readable, and follows language-specific best practices.
- Meaningful variable/function names and comments explain key logic.
- Modular design (e.g., separate hash function, helper functions).

### Edge Case Handling (15%)
- Solutions handle edge cases gracefully:
  - Exercise 1: Missing/out-of-order fragments, tampered hashes.
  - Exercise 2: Single center, zero fragments, high risk factors.
  - Exercise 3: No path, no compression nodes, cyclic graphs.
- Error messages or safe defaults for invalid inputs.

### Documentation (10%)
- README is clear, with complete run instructions and assumptions.
- Brief, accurate explanations of each solution’s approach.

## Submission Deadline
- The 3-hour time limit begins when you receive the challenge.
- Submit the repository URL to evaluation team within 3 hours of receipt.
- Late submissions will not be accepted unless pre-approved due to extenuating circumstances.

## Questions?
Contact the evaluation team for clarifications or to schedule receipt of the challenge. We encourage you to confirm the programming language and any setup questions before starting.

Good luck, and we look forward to reviewing your solutions in this cyberpunk challenge!

