# The feature planning doc

This document is to plan out what features I need to build with this cli application. This application will help build out test plans based on a template and type of validation.

## Features To Build out
* Build charm form that collects the inputs - **Complete**
  - Ticket
  - Jira Link
  - Type of validation required for testplan
  - Automation required
* Build template to dynamically update info - **Complete**
  - Use go lib text/template
* Build write functionality **Complete**
  - to write to specified path **Complete**
  - Update form to collect path field to write to **Complete**

## Future state features
* Build template for Jira ticket creation
* Automation write to Jira
* Live build of form inputs as TUI
* add Test Validation section in form to build out some validation steps
* New Feature / bug type ticket category
  - Dynamic form inputs based on type ex: repro steps for bugs
