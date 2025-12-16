**Title**
LinkedIn Automation Proof of Concept (Go + Rod)

**Overview**
This project is a Go-based proof of concept that demonstrates browser automation using the Rod library.
It focuses on clean architecture, human-like behavior simulation, and basic anti-detection strategies.

This project is for educational and technical evaluation purposes only.

Tech Stack

Go

Rod (browser automation)

Chromium

GitHub

Project Structure

main.go – application entry point
internal/browser – browser startup and control
internal/auth – login simulation
internal/search – profile search simulation
internal/connect – connection request simulation
internal/message – messaging simulation
internal/stealth – human-like behavior and anti-detection logic
internal/config – configuration management

**Setup Instructions**
Clone repository
git clone https://github.com/YOUR_USERNAME/linkedin-assignment.git

cd linkedin-assignment

Install dependencies
go mod tidy

Create environment file
Copy .env.example to .env and fill values.

Run application
go run .

Environment Variables

LINKEDIN_EMAIL
LinkedIn login email

LINKEDIN_PASSWORD
LinkedIn login password

DAILY_LIMIT
Maximum number of daily actions

Implemented Stealth Techniques

Randomized delays between actions

Human-like mouse movement

Human-like typing simulation

Natural scrolling behavior

Rate limiting for actions

Browser fingerprint masking (basic)

Demonstration Video
Video walkthrough link:


This project is created strictly for educational and technical demonstration purposes.
Automating LinkedIn violates LinkedIn Terms of Service.
Do not use this tool on real accounts or in production environments.
