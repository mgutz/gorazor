'use strict';
var gulp = require('gulp');
var sh = require('gee-shell');
var _ = require('lodash');

var watches = {
  views: {
    watch: ['views/**/*.gohtml']
  }
};

gulp.task('views', function(done) {
  sh.run('gorazor views views', done);
});

gulp.task('watch', Object.keys(watches), function(done) {
  _.forOwn(watches, function(val, key) {
    gulp.watch(val.watch || val.src, [key]);
  });
});
