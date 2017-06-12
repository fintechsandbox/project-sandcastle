![FinTech Sandbox](http://fintechsandbox.org/sites/default/files/fintech-logo_0_0.png)
FinTech Sandbox Members Only
============================
This repository and the associated Wiki are to be used only by active members of the FinTech Sandbox Community.
We are calling this resource **Project Sandcastle**, and we have designed it to facilitate the quick and
effective adoption of a variety of financial data services.

# Project Sandcastle
http://www.fintechsandbox.org

This resource is organized into the following sections:

* **Wiki**
  * [Sandbox Members](https://github.com/closedLoop/fintech-sandbox-curation/wiki/Sandbox-Members)
    * Contains a page for each Current Sandbox Member and Alumnus in order to maintain appropriate contact information
  * [Data Providers](https://github.com/closedLoop/fintech-sandbox-curation/wiki/Data-Providers)
    * Go here to **request introductions** to data providers
    * Submit and view reviews for each data provider
    * View Product Descriptions and Features
      * Attach relevant **documentation** and **marketing materials**
* **Code Base**
  * [Data Provider Code Base](./data_providers)
    * Each Data provider folder has three sub-folders:
      * **API​**: Documentation that references product descriptions in the Wiki
      * **Issues​**: Solutions to known bugs and references to solutions on StackOverflow
      * **Integrations​**: Solutions to convert external data sources into useful formats for databases, visualizations, etc...
  * [Member's Sandbox](./member_sandbox)
    * A collection of data-provider-agnostic solutions to pressing problems for members in the Sandbox Community
    * Be creative and collaborative here!
  * [FinTech Sandbox Scripts](./fintech_sandbox)
    * Code required to generate specific subsets of this [codebase](https://github.com/closedLoop/fintech-sandbox-curation) and [Wiki](https://github.com/closedLoop/fintech-sandbox-curation/wiki)
* **Slack**
  * Some problems we face in launching a startup require **communication, not code**.  For this we have Slack: [FinTechSandbox.slack.com](https://fintechsandbox.slack.com)
  * Go there to share your company updates, ask for help and help another entrepreneur / data-geek out.

## Contribution Guidelines
 - There is a lot of good resources here and the FinTech Sandbox team is working hard to create a great environment for you all to succeed.
 - Please read [CONTRIBUTING.md](./CONTRIBUTING.md) to understand how to keep this place organized and useful to others

### **TL;DR** Contributing ###
 1. Keep reviews and descriptions of data providers in the Wiki
 1. Keep contact information for your company up-to-date in the Wiki
 1. Code should be moderately documented and minimally functional with a clear description of the problem it is trying to address
 1. All programming languages are welcome
 1. Programs that rely on multiple files should be placed in their own sub-folder
 1. Follow the templates provided:  
   1. Wiki:  [New Participant](https://github.com/fintechsandbox/project-sandcastle/wiki/member_template), [New Data Provider](https://github.com/fintechsandbox/project-sandcastle/wiki/provider_template)
   1. Code Base: [New Data Provider](./data_providers/__PROVIDER_TEMPLATE__)
 
 ### Note to developers
 This repo uses Git Large File System: ``git lfs`` so that it more effectively handles large documentation files.  You may have trouble cloning this repo unless you install it.  Look for instructions here:  [https://git-lfs.github.com/](https://git-lfs.github.com/)
 
