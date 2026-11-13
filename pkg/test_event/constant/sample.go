package constant

const FixtureMixedStream = `go: downloading example.com/module v1.0.0
{"Time":"2026-11-13T17:06:33.875056+01:00","Action":"start","Package":"probe/good"}
{"Time":"2026-11-13T17:06:34.041646+01:00","Action":"run","Package":"probe/good","Test":"TestPasses"}
{"Time":"2026-11-13T17:06:34.041749+01:00","Action":"pass","Package":"probe/good","Test":"TestPasses","Elapsed":0}
{"Time":"2026-11-13T17:06:34.041757+01:00","Action":"run","Package":"probe/good","Test":"TestFails"}
{"Time":"2026-11-13T17:06:34.041765+01:00","Action":"output","Package":"probe/good","Test":"TestFails","Output":"    good_test.go:7: want 1, got 2\n","OutputType":"error"}
{"Time":"2026-11-13T17:06:34.041775+01:00","Action":"fail","Package":"probe/good","Test":"TestFails","Elapsed":0}
{"Time":"2026-11-13T17:06:34.041779+01:00","Action":"run","Package":"probe/good","Test":"TestPanics"}
{"Time":"2026-11-13T17:06:34.044238+01:00","Action":"output","Package":"probe/good","Test":"TestPanics","Output":"panic: boom [recovered, repanicked]\n"}
{"Time":"2026-11-13T17:06:34.044819+01:00","Action":"fail","Package":"probe/good","Test":"TestPanics","Elapsed":0}
{"Time":"2026-11-13T17:06:34.044831+01:00","Action":"fail","Package":"probe/good","Elapsed":0.17}
`

const FixtureBuildFailureStream = `{"ImportPath":"probe/broken [probe/broken.test]","Action":"build-output","Output":"# probe/broken [probe/broken.test]\n"}
{"ImportPath":"probe/broken [probe/broken.test]","Action":"build-output","Output":"broken/broken.go:4:2: declared and not used: x\n"}
{"ImportPath":"probe/broken [probe/broken.test]","Action":"build-fail"}
{"Time":"2026-11-13T17:06:34.24902+01:00","Action":"start","Package":"probe/broken"}
{"Time":"2026-11-13T17:06:34.249081+01:00","Action":"output","Package":"probe/broken","Output":"FAIL\tprobe/broken [build failed]\n","OutputType":"frame"}
{"Time":"2026-11-13T17:06:34.249096+01:00","Action":"fail","Package":"probe/broken","Elapsed":0,"FailedBuild":"probe/broken [probe/broken.test]"}
`

const FixtureSetupPanicStream = `{"Time":"2026-11-13T17:12:46.983052+01:00","Action":"start","Package":"probe/setup"}
{"Time":"2026-11-13T17:12:47.216128+01:00","Action":"output","Package":"probe/setup","Output":"panic: fixture missing\n"}
{"Time":"2026-11-13T17:12:47.216559+01:00","Action":"output","Package":"probe/setup","Output":"FAIL\tprobe/setup\t0.233s\n","OutputType":"frame"}
{"Time":"2026-11-13T17:12:47.216581+01:00","Action":"fail","Package":"probe/setup","Elapsed":0.234}
`

const FixtureNoTestsStream = `{"Time":"2026-11-13T17:06:34.505209+01:00","Action":"start","Package":"probe/good"}
{"Time":"2026-11-13T17:06:34.570947+01:00","Action":"output","Package":"probe/good","Output":"testing: warning: no tests to run\n"}
{"Time":"2026-11-13T17:06:34.571034+01:00","Action":"output","Package":"probe/good","Output":"PASS\n","OutputType":"frame"}
{"Time":"2026-11-13T17:06:34.571272+01:00","Action":"output","Package":"probe/good","Output":"ok  \tprobe/good\t0.066s [no tests to run]\n"}
{"Time":"2026-11-13T17:06:34.571286+01:00","Action":"pass","Package":"probe/good","Elapsed":0.066}
`
