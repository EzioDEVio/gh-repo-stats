# gh-repo-stats

## Overview

`gh-repo-stats` is a GitHub CLI extension designed and developed by Mohammed Aldaraji: (https://github.com/theslash84) to streamline the process of fetching and displaying statistics for any GitHub repository. As the sole developer of this extension, I aimed to create a tool that enhances the GitHub CLI's functionality, making it easier for users to get quick insights into repository metrics.

## Key Features and Benefits

- **Quick Access to Repository Stats**: Retrieve essential statistics such as star count, fork count, and open issues with a single command.
- **Enhance Productivity**: Save time and effort by avoiding manual repository checks or navigating through GitHub's UI.
- **Integration with GitHub CLI**: Seamlessly extends the GitHub CLI, maintaining a consistent command-line experience.
- **Open Source**: Freely available for use and modification, fostering community collaboration and improvement.

## Installation

To install `gh-repo-stats`, you need to have the GitHub CLI installed on your machine. If you haven't installed it yet, follow the [official installation guide](https://cli.github.com/manual/installation). Once you have the GitHub CLI, install the extension using the following command:

```shell
gh extension install theslash84/gh-repo-stats
```

---

## Usage
After installation, you can fetch repository statistics by executin
```
gh repo-stats -owner <owner> -repo <repository name>
```
For example, to fetch statistics for the gh-repo-stats repository itself:
```
gh repo-stats -owner theslash84 -repo gh-repo-stats

```
This command will display the number of stars, forks, and open issues for the specified repository.

Contributing
As the project's sole developer, I welcome contributions that can enhance gh-repo-stats. Whether it's feature suggestions, bug reports, or code contributions, feel free to share your input. Please submit your contributions through issues and pull requests.

License
gh-repo-stats is distributed under the MIT License. See the LICENSE file for more details.

Sharing with GitHub Extension Library
To share gh-repo-stats with the GitHub extension library and make it discoverable for other users, you can add the gh-extension topic to your repository:

Navigate to your repository on GitHub.
Click on the "About" section on the right-hand side.
Add gh-extension to the "Topics" field.
By adding this topic, your extension becomes part of the broader collection of GitHub CLI extensions, making it discoverable for users interested in extending their CLI capabilities.




