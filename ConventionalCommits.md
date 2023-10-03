# Conventional Commits Notes

Commit Message Structure

Each commit message consists of a header, a body and a footer. The header has a special format that includes a type, a 
scope and a subject:

```text
<type>(<scope>): <subject>
<BLANK LINE>
[optional body]
<BLANK LINE>
[optional footer(s)]
```

The header is mandatory and the scope of the header is optional.

Any line of the commit message cannot be longer than 100 characters! This allows the message to be easier to read on
GitHub as well as in various git tools.

The footer should contain a closing reference to an issue if any.

Question: What the heck does this mean?

Types:
- fix: A bug fix
- feat: A new feature
- build: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
- chore
- ci: Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)
- docs: Documentation only changes
- style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- refactor: A code change that neither fixes a bug nor adds a feature
- perf: A code change that improves performance.
- test: Adding missing tests or correcting existing tests
- revert: If a commit revers a previous commit, it should begin with `revert:`, followed by the header of the reverted commit.  In the body it should say `This reverts commit <hash>.`, where the hash is the SHA of the commit being reverted.

## Scope

The scope should be the name of the package affected (as perceived by the person reading the changelog generated from commit messages.)

Supported scopes:

- animations
- common
- compiler
- compiler-cli
- core
- elements
- forms
- http
- language-service
- platform-browser
- platform-browser-dynamic
- platform-server
- platform-webworker
- platform-webworker-dynamic
- router
- upgrade

## Subject

The subject contains a succinct description of the change:

- use the imperative, present tense: "change" not "changed" nor "changes"
- don't capitalize the first letter
- no dot (.) at the end

## Body

Just as in the subject, use the imperative, present tense: "change" not "changed" nor "changes". The body should include

## Footer

The footer should contain any information about Breaking Changes and is also the place to reference GitHub issues that this commit Closes.

Breaking Changes should start with the word BREAKING CHANGE: with a space or two newlines. The rest of the commit message is then used for this.

A detailed explanation can be found in this document.

## References

- [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
- [Commit Message Guidelines](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#-commit-message-guidelines)