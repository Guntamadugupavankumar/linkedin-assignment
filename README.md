LinkedIn Automation Assignment
Go + Rod Proof of Concept

**Overview**
This project is a Go-based proof of concept that demonstrates advanced browser automation using the Rod library.
It focuses on clean architecture, human-like behavior simulation, and anti-bot stealth techniques.
This is built strictly for technical evaluation and learning.

Tech Stack

Go

Rod browser automation library

Chrome browser

Environment variables for configuration

**Project Structure**

main.go
Application entry point and execution flow

internal/browser
Browser startup and lifecycle handling

internal/auth
Login flow simulation

internal/search
Profile search simulation

internal/connect
Connection request simulation

internal/message
Messaging simulation

internal/stealth
Anti-bot and human-like behavior techniques

internal/config
Environment-based configuration loading

**Core Features**

Modular Go architecture

Browser automation using Rod

Environment-based configuration

Rate-limited actions

Clear execution flow and logs

Implemented Stealth Techniques

Randomized delays between actions

Human-like mouse movement paths

Browser fingerprint masking

Human-like typing simulation

Human-like scrolling behavior

Rate limiting and daily action limits

These techniques simulate realistic user behavior and reduce automation signatures.

**Configuration**
This project uses environment variables.

Create a local .env file using the provided example.

.env.example
LINKEDIN_EMAIL=your_email@example.com

LINKEDIN_PASSWORD=your_password_here
DAILY_LIMIT=5

Do not commit real credentials.
Only .env.example should be pushed to GitHub.

**How to Run**

Install Go

Install Google Chrome

Clone the repository

Create a .env file locally

Run the application

**Command**
go run .

Expected Output

App started

Config loaded

Browser started

Stealth actions executed

Simulated login, search, connect, and message flows

App finished

Disclaimer
Educational purpose only.
This project is designed for technical evaluation and demonstration.
Automating LinkedIn violates their Terms of Service.
Do not use this tool on real accounts.
Do not deploy in production environments.

Author
Pavan Kumar
LinkedIn Automation Technical Assignment

**GOOGLE DRIVE LINK**
https://drive.google.com/file/d/1ENnzzEr9hS_0Nq7OW2OxBVbUzMRIYeUD/view?usp=sharing
