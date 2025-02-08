# Test plan maker

## About
In my day-to-day work, I find myself manually typing out test plans. I prefer a specific structure and found myself repeatedly typing out these fields. I decided to build this CLI tool to generate a test plan and write it to a specified path, or to the user's Desktop if no path is provided. I hope you find it as useful in building out your test plans as I have.

***

## How To Use
To build a new test plan run:
```
BuildTestPlan(ver#) *the version number is part of the release you have downloaded
```
example: BuildTestPlan1

![Demo](DemoV1.0.1.gif)

### Fields

* User Story Details Section
    * Test Plan title: used as the title of the markdown file generated
    * User Story: The User Story ID (Ex. TST-100 will be used to generate the md file)
    * User Story Link: will generate a link to the user story tied to the test plan
    * Product Manager/ Product Owner
    * Software Engineer
    * Quality Assurance
    * User Story Components: Components that will be deployed

* Test Plan Details Section:
    * Validation Type: sigle toggle to select type of validation involved with the user story
        * Current Options:
            1. UI Validation
            2. API Validation
            3. UI & API Validation
    * Is Automation Needed: Does this user story require automation to be built out
    * Description: A description of the work for the user story and any necessary details for the test plan
* Save file to path: File path you wish to save the file to. If left blank, it defaults to the users Desktop.
    * Saved as md file: ``` UserStory123.md ```
* Write Test Plan Confirmation: Once you highlight yes and submit the builder will generate a Markdown file to the specified or default path with the naming convention mentioned above.

***

## Prerequisites

If you want to run this project locally you need to install go. If you don't have go installed on your machine you can download it from https://go.dev/dl/


### To Run without compiling
To run locally from within the DIR this project is in input into terminal ``` Make run ```

### Compilation

In a terminal window in the directory this project is based run ```Make compile``` this will compile a binary based on your machine's OS with the name **BuildTestPlan**.

### Read

For this initial phase and the demo in include in this readme, I use the library glow from charmbracelet to read the md file I just created. I figured I'd add it to the make file commands in case you want to quickly view the .md file in your terminal as well!

**To execute**
Make sure you have *glow* installed on your machine [Installation Instruction](https://github.com/charmbracelet/glow#package-manager)

In terminal window in directory this project is based run:

```Make read filePath=filepath goes here```

Example:
```Make read filePath=./testSaveFileDir/example.md```
