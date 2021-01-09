# [badges.joshuamiller.net](https://badges.joshuamiller.net/)

[![Go Report Card](https://goreportcard.com/badge/github.com/JoshuaTheMiller/badges.joshuamiller.net)](https://goreportcard.com/report/github.com/JoshuaTheMiller/badges.joshuamiller.net)

A [Vercel](https://vercel.com/) Project containing Functions to support custom [Shields](https://shields.io/) for my various repositories. An article will most likely be coming out soon, so follow me on Medium if you want a notification about it!

[![Medium](https://img.shields.io/badge/Follow%20on-Medium-lightgrey)](https://medium.com/@JoshuaTheMiller)

## Examples

The following are some of my favorite custom badges, and a quick explanation for their existence.

### Application Destroyed Badge

[![Custom badge](https://img.shields.io/endpoint?style=flat-square&url=https%3A%2F%2Fbadges.joshuamiller.net%2Fapi%2Fshield%2Fcodedeployexample.go)](https://dev.azure.com/JoshuaTheMiller/PublicExamples/_build?definitionId=3)

Within my [Experiments repository](https://github.com/JoshuaTheMiller/Experiments) is an [example of using AWS CodeDeploy](https://github.com/JoshuaTheMiller/Experiments/tree/main/aws/codedeploy). This example, when fully deployed, creates a website that folks can look at. However, to save cost the website is torn down automatically after some time. To help readers of the README determine the status of the website, I added a badge that would communicate whether the site was AVAILABLE or if it was DESTROYED.