<p align="center">
  <img
    alt="toolbox-export"
    src="https://shieldcn.dev/header/gradient.svg?title=toolbox-export&subtitle=Making+life+easier+working+with+toolbox+containers&mode=dark"
  />
</p>

<p align="center">
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/stars/Pratyay360/toolbox-export.svg?variant=secondary" /><img alt="GitHub Stars" src="https://shieldcn.dev/github/stars/Pratyay360/toolbox-export.svg?variant=secondary&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/forks/Pratyay360/toolbox-export.svg?variant=secondary" /><img alt="GitHub Forks" src="https://shieldcn.dev/github/forks/Pratyay360/toolbox-export.svg?variant=secondary&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/watchers/Pratyay360/toolbox-export.svg?variant=secondary" /><img alt="Watchers" src="https://shieldcn.dev/github/watchers/Pratyay360/toolbox-export.svg?variant=secondary&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/branches/Pratyay360/toolbox-export.svg?variant=ghost" /><img alt="Branches" src="https://shieldcn.dev/github/branches/Pratyay360/toolbox-export.svg?variant=ghost&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/contributors/Pratyay360/toolbox-export.svg?theme=emerald" /><img alt="Contributors" src="https://shieldcn.dev/github/contributors/Pratyay360/toolbox-export.svg?theme=emerald&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/last-commit/Pratyay360/toolbox-export.svg?variant=secondary" /><img alt="Last commit" src="https://shieldcn.dev/github/last-commit/Pratyay360/toolbox-export.svg?variant=secondary&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/commits/Pratyay360/toolbox-export.svg?variant=secondary" /><img alt="Commits" src="https://shieldcn.dev/github/commits/Pratyay360/toolbox-export.svg?variant=secondary&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/open-issues/Pratyay360/toolbox-export.svg?variant=secondary" /><img alt="Open issues" src="https://shieldcn.dev/github/open-issues/Pratyay360/toolbox-export.svg?variant=secondary&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/closed-issues/Pratyay360/toolbox-export.svg?variant=ghost" /><img alt="Closed issues" src="https://shieldcn.dev/github/closed-issues/Pratyay360/toolbox-export.svg?variant=ghost&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/open-prs/Pratyay360/toolbox-export.svg?variant=secondary" /><img alt="Open PRs" src="https://shieldcn.dev/github/open-prs/Pratyay360/toolbox-export.svg?variant=secondary&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/closed-prs/Pratyay360/toolbox-export.svg?variant=ghost" /><img alt="Closed PRs" src="https://shieldcn.dev/github/closed-prs/Pratyay360/toolbox-export.svg?variant=ghost&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/merged-prs/Pratyay360/toolbox-export.svg?variant=ghost" /><img alt="Merged PRs" src="https://shieldcn.dev/github/merged-prs/Pratyay360/toolbox-export.svg?variant=ghost&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/release/Pratyay360/toolbox-export.svg" /><img alt="Release" src="https://shieldcn.dev/github/release/Pratyay360/toolbox-export.svg?mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/ci/Pratyay360/toolbox-export.svg?variant=secondary" /><img alt="CI" src="https://shieldcn.dev/github/ci/Pratyay360/toolbox-export.svg?variant=secondary&mode=light" /></picture>
  <picture><source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/github/license/Pratyay360/toolbox-export.svg?variant=ghost" /><img alt="License" src="https://shieldcn.dev/github/license/Pratyay360/toolbox-export.svg?variant=ghost&mode=light" /></picture>
</p>

## Overview

toolbox-export makes working with applications installed inside a Toolbox container easier. Export desktop launchers, icons, and command-line wrappers while keeping applications in the container.

## Installation

Download and extract the Linux archive for your architecture from [GitHub Releases](https://github.com/Pratyay360/toolbox-export/releases):

```bash
curl -fSsL https://raw.githubusercontent.com/Pratyay360/toolbox-export/refs/heads/main/install.sh | sh
```

```bash
mise github:Pratyay360/toolbox-export@latest
go install github.com/Pratyay360/toolbox-export@latest
stew s Pratyay360/toolbox-export
mise install go:github.com/Pratyay360/toolbox-export@latest
curl -sf http://goblin.run/github.com/Pratyay360/toolbox-export | PREFIX=~/.local/bin sh
```

## Usage

Toolbox-export is a batteries-included utility. It performs all the necessary actions on your behalf.

```bash
toolbox-export init (initializes toolbox-export in your home directory)
toolbox-export <app name> (creates both binary and desktop entry wrapper; automatically manages all the necessary cases, like a CLI-only application or a GUI application)

toolbox-export -r / --rm <appname> (removes both binary and desktop entry wrapper)
```

Toolbox containers run the application in a sandboxed environment. This is one of the most developer friendly ways to run and test applications.

## License

[Apache License 2.0](LICENSE).
