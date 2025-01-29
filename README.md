# Test plan maker

## About
In my day-to-day work, I find myself manually typing out test plans. I prefer a specific structure and found myself repeatedly typing out these fields. I decided to build this CLI tool to generate a test plan and write it to a specified path, or to the user's Desktop if no path is provided. I hope you find it as useful in building out your test plans as I have.

## How To
To build a new test plan run:
```
MakeTestPlan(ver#)
```
example: MakeTestPlan.1.0

### Fields

* Test Plan title: used as the title of the markdown file generated
* User Story Link: will generate a link to the user story tied to the test plan
* Validation Type: sigle toggle to select type of validation involved with the user story
  - Current Options:
    1. UI Validation
    2. API Validation
    3. UI & API Validation
* Description: A description of the work for the user story and any necessary details for the test plan
* Save file to path: File path you wish to save the file to. If left blank, it defaults to the users Desktop.
  - Saved as md file: ``` UserStory123.md ```
* Is Automation Needed: Does this user story require automation to be built out
