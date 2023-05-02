# Inorder to apply pre-commit-hooks:

Install pre-commit package manager

- ```pip install pre-commit ```

Run pre-commit install to set up the git hook scripts

- ```pre-commit install```

Install gofumpt golang linter

- ```go install mvdan.cc/gofumpt@latest ```

Migrate list configuration to new map configuration.

- ```pre-commit migrate-config```

it's usually a good idea to run the hooks against all of the files when adding new hooks (usually pre-commit will only run on the changed files during git hooks)

- ```pre-commit run --all-files```