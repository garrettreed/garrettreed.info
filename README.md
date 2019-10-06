# Static Site Template

-   Babel transpilation with ES2015 defaults
-   Live reloading with html injection
-   Sass
-   `.env` environment configuration
-   Sourcemaps (in dev env only)
-   Templates, partials, data sources with Mustache
-   Polyfills and css autoprefixing for default browser support. Autoprefixer and Babel use [browserslist]https://github.com/browserslist/browserslist#queries) to configure browser support. Default is `> 0.5%, last 2 versions, Firefox ESR, not dead`, but this can be configured.

## Notes

-   `*.html` files placed in `public/` will be ignored, as will any file placed in `public/dist/`. You can add html files to `templates/` and they will be copied over in the build step.

## TODO

-   Image task
-   Fonts?
-   Module bundler?
-   Favicons?
