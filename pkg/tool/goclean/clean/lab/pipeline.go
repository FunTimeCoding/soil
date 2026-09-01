package lab

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/option"
	"log"
)

func Pipeline(
	o *option.Clean,
	c *gitlab.Client,
	p *project.Project,
) {
	branches := c.MustBranches(p.Identifier)
	mainBranch := branch.MainBranch(branches)
	mainHash := mainBranch.Raw.Commit.ID

	if mainHash == "" {
		log.Panic("empty branch hash")
	}

	if o.Verbose {
		console.Format("Default branch: %s\n", p.Raw.DefaultBranch)
		console.Format("Main branch: %s\n", mainBranch.Name)

		for _, b := range branches {
			console.Format("Branch: %s %s\n", b.Name, b.Raw.Commit.ID)
		}

		console.Format("Main hash: %s\n", mainHash)
	}

	pipelines := c.MustPipelines(p.Identifier)

	if len(pipelines) == 0 {
		return
	}

	latestSemantic := pipeline.LatestSemantic(pipelines)
	latestMain := pipeline.LatestMain(pipelines, mainHash)

	if latestSemantic == nil && o.Verbose {
		errors.Warning("no latest semantic pipeline found")
	}

	if latestMain == nil && o.Verbose {
		errors.Warning("no latest main pipeline found")
	}

	for _, i := range pipelines {
		if latestSemantic != nil &&
			i.Ref == latestSemantic.Ref &&
			i.SHA == mainHash {
			if o.Verbose {
				console.Format("Skip pipeline (sematic): %s %s\n", i.Ref, i.SHA)
			}

			continue
		}

		if latestMain != nil &&
			i.Ref == mainBranch.Name &&
			i.SHA == latestMain.SHA {
			if o.Verbose {
				console.Format("Skip pipeline (main): %s %s\n", i.Ref, i.SHA)
			}

			continue
		}

		if o.Verbose {
			console.Format("Pipeline: %s %s\n", i.Ref, i.SHA)
		} else {
			console.Format("Pipeline: %s\n", i.Ref)
		}

		c.MustDeletePipeline(p.Identifier, i.ID)
	}
}
