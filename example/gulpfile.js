'use strict';
var gulp = require('gulp');
var sh = require('gee-shell');
var _ = require('lodash');
var less = require('gulp-less');

var sources = {
  views: {
    watch: ['views/**/*.gohtml']
  },

  // {{version}} is for proper cache-busting.  Simply search and replace
  // "{{version}}" in all files when deploying, rename {{version}} directory
  // on your CDN and profit.
  styles: {
    src: ['public/{{version}}/css/style.less'],
    watch: ['public/{{version}}/css/**/*.less']
  }
};

gulp.task('styles', function() {
  gulp.src(sources.styles.src)
    .pipe(less())
    .pipe(gulp.dest('dist/{{version}}/css'));
});

gulp.task('views', function(done) {
  sh.run('gorazor views views', done);
});

gulp.task('all', Object.keys(sources));

gulp.task('watch', Object.keys(sources), function(done) {
  _.forOwn(sources, function(val, key) {
    gulp.watch(val.watch || val.src, [key]);
  });
});
