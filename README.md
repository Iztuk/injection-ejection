# Injection-Ejection
A demo app that showcases how SQL Injection vulnerabilities are introduced and how to fix them using secure coding practices.

---

## Scenario: The Curious Manager at Injec Corp

This project simulates a vulnerable internal employee portal for a fictional company called **Injec Corp**, where managers use a simple search form to look up employee records by name.

The search functionality was implemented quickly — and carelessly — using raw SQL queries with direct string interpolation. This allowed attackers to craft malicious input that could:

- Return **all employee records** regardless of search filters (`' OR 1=1--`)
- Cause **unexpected behavior or errors**
- In some cases, even execute **destructive queries** (e.g., `Robert'); DROP TABLE Employees;--`)

---

## Objectives

This project demonstrates:
1. **How SQL injection vulnerabilities occur** in real-world code.
2. **How to safely mitigate them** using parameterized queries and input validation.
2. The critical role of **secure coding practices** in preventing injection attacks - one of the [OWASP Top 10](https://owasp.org/Top10/A03_2021-Injection/)

---

## Features
- Employee search form
- Vulnerable endpoints: `/vulnerable`
- Secure endpoints: `/secure`
- Sample SQL queries and inputs to demonstrate injection vectors

---

> *This project focuses solely on SQL Injection vulnearbilites, from exploitation to defense, as a hands-on learning and demonstration tool for secure software engineering. If there are any suggestions, please leave contact me through my email on my GitHub profile!*