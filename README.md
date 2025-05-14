# Injection-Ejection
A demo app that showcases how SQL Injection vulnerabilities are introduced and how to fix them using secure coding practices.

---

## Scenario

This project simulates a vulnerable internal employee portal for a company called **Injec Corp**.

The app inclueds a search feature that was originally implemented without input sanitization, allowing attackers to perform SQL Injection attacks (e.g., `', OR 1=1--` or even `Robert'); DROP TABLE Employees;--`).

The goal of this project is to demonstrate:
1. How SQL injection works through hands-on exploitation.
2. How to mitigate it using prepared statements and input sanitization/validation.
3. The importance of secure coding practices in real-world applications.

---

## Features
- Landing page
- Employee search by name (vulnerable to injection)
- Toggle between `/vulnerable` and `/secure` endpoints