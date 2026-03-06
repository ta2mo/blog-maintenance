# AGENTS.md

## Project-specific rules

- This project does not have a `docs/` directory. Do not assume `docs/` exists when gathering requirements/specifications.
- Blog post Vue files under `nuxt-template/pages/post/` are generated from markdown files.
- To regenerate post pages, run `./blog-maintenance convert` from the repository root.
- For tasks involving `/post` regeneration, the user will run the regeneration command unless explicitly requested otherwise.
