<p align="center">
  <img
    alt="toolbox-export"
    src="https://shieldcn.dev/header/gradient.svg?title=toolbox-export&subtitle=Improved+usability+for+applications+installed+inside+toolbx+containers&mode=dark"
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

`toolbox-export` makes applications installed inside a [Toolbx](https://containertoolbx.org/) container available on your Linux host. Export desktop launchers, icons, and command-line wrappers while keeping applications in the container.

## Installation

Download and extract the Linux archive for your architecture from [GitHub Releases](https://github.com/Pratyay360/toolbox-export/releases):

```bash
curl -sf http://goblin.run/github.com/Pratyay360/toolbox-export | PREFIX=~/.local/bin sh
```

```bash
mise github:Pratyay360/toolbox-export@latest
```

```bash
go install github.com/Pratyay360/toolbox-export@latest
```

```bash
stew s Pratyay360/toolbox-export
```

```bash
mise install go:github.com/Pratyay360/toolbox-export@latest
```

```bash
curl -fSsL https://raw.githubusercontent.com/Pratyay360/toolbox-export/refs/heads/main/install.sh | sh



eget Pratyay360/toolbox-export
```

### Usage

run

```sh
toolbox-export init
```

to let the tool setup things for you. it's performing some actions on your behalf.

creating a directory $$HOME/.local/toolbox$ if not exist and add that to `PATH` variable. Some operations can be performed like.

```sh
toolbox-export <app name> (creates both binary and desktop entry wrapper)
```

## License

[Apache License 2.0](LICENSE).
