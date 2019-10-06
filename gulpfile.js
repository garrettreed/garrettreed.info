require("dotenv").config();
const gulp = require("gulp");
const browserSync = require("browser-sync").create();
const sass = require("gulp-sass");
const babel = require("gulp-babel");
const uglify = require("gulp-uglify");
const sourcemaps = require("gulp-sourcemaps");
const autoprefixer = require("gulp-autoprefixer");
const mustache = require("gulp-mustache");
const concat = require("gulp-concat");
const gulpif = require("gulp-if");
const rev = require("gulp-rev");
const clean = require("gulp-clean");
const revRewrite = require("gulp-rev-rewrite");

const isProdEnv = process.env.ENVIRONMENT === "production";
console.log(`Acting on ${process.env.ENVIRONMENT} environment.`);

const distPath = "public/dist/";
const paths = {
    styles: {
        src: ["libs/**/*.css", "styles/**/*.scss"],
        dest: distPath
    },
    scripts: {
        src: ["libs/**/*.js", "js/**/*.js"],
        dest: distPath
    },
    templates: {
        src: ["templates/**/*.mustache", "templates/**/*.html", "!templates/partials/**"],
        dest: "public/"
    }
};

function styles() {
    return gulp
        .src(paths.styles.src)
        .pipe(sourcemaps.init())
        .pipe(sass({ outputStyle: "compressed" }).on("error", sass.logError))
        .pipe(autoprefixer())
        .pipe(concat("dist.css"))
        .pipe(gulpif(!isProdEnv, sourcemaps.write(".")))
        .pipe(gulp.dest(paths.styles.dest))
        .pipe(browserSync.stream({ match: "**/*.css" }));
}

function scripts() {
    return gulp
        .src(paths.scripts.src)
        .pipe(sourcemaps.init({ loadMaps: true }))
        .pipe(concat("dist.js"))
        .pipe(babel({ presets: ["@babel/preset-env"] }))
        .pipe(uglify())
        .on("error", swallowError)
        .pipe(gulpif(!isProdEnv, sourcemaps.write(".")))
        .pipe(gulp.dest(paths.scripts.dest));
}

function templates() {
    return gulp
        .src(paths.templates.src)
        .pipe(mustache("data.json", { extension: ".html" }, {}))
        .pipe(gulp.dest(paths.templates.dest));
}

function buildManifest() {
    return gulp
        .src(`${distPath}/*.{js,css}`)
        .pipe(rev())
        .pipe(gulp.dest(distPath))
        .pipe(
            rev.manifest(`${distPath}/rev-manifest.json`, {
                base: distPath,
                merge: true
            })
        )
        .pipe(gulp.dest(distPath));
}

function writeManifest() {
    return gulp
        .src(`${paths.templates.dest}**/*.html`)
        .pipe(revRewrite({ manifest: gulp.src("public/dist/rev-manifest.json") }))
        .pipe(gulp.dest(paths.templates.dest));
}

function cleanDist() {
    return gulp.src(`${distPath}/*.{js,css,map}`, { read: false }).pipe(clean());
}

function watch() {
    gulp.watch(paths.scripts.src, scripts).on("change", browserSync.reload);
    gulp.watch(paths.styles.src, styles);
    gulp.watch("./data.json", templates).on("change", browserSync.reload);
    gulp.watch(paths.templates.src, templates).on("change", browserSync.reload);
}

function swallowError(error) {
    console.log(error.toString());
    this.emit("end");
}

function serve() {
    browserSync.init({
        server: "public"
    });
    watch();
}

const build = gulp.series(
    cleanDist,
    gulp.parallel(styles, scripts, templates),
    buildManifest,
    writeManifest
);

const dev = gulp.series(cleanDist, gulp.parallel(styles, scripts, templates), serve);

exports.styles = styles;
exports.scripts = scripts;
exports.templates = templates;
exports.watch = watch;
exports.build = build;
exports.dev = dev;
exports.default = dev;
