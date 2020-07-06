# goalg

[![Build Status](https://travis-ci.org/vogo/goalg.png?branch=master)](https://travis-ci.org/vogo/goalg)
[![codecov](https://codecov.io/gh/vogo/goalg/branch/master/graph/badge.svg)](https://codecov.io/gh/vogo/goalg)
[![GoDoc](https://godoc.org/github.com/vogo/goalg?status.svg)](https://godoc.org/github.com/vogo/goalg)
[![Go Report Card](https://goreportcard.com/badge/github.com/vogo/goalg)](https://goreportcard.com/report/github.com/vogo/goalg)
![license](https://img.shields.io/badge/license-Apache--2.0-green.svg)

## 二分查找
   
**核心思想**: 
- 给定排序的数组，选定中间位置，和目标值比较；
- 如果小于目标值，则在左侧继续查找；
- 如果大于目标值，则在右侧继续查找。

**算法关键**: 
- mid在边界情况时取值: `r = l + 1` 情况下, 即左右指针相邻的情况, `mid = l = l + (r - l)/2`,  `mid = r = l + (r - l + 1)/2`;
- 避免死循环: 确保以上边界情况 l 或 r 的值变化后下一次循环mid取得不同的值，或能跳出循环。
