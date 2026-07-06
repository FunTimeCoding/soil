package lab

import (
	"fmt"
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
		fmt.Printf("Default branch: %s\n", p.Raw.DefaultBranch)
		fmt.Printf("Main branch: %s\n", mainBranch.Name)

		for _, b := range branches {
			fmt.Printf("Branch: %s %s\n", b.Name, b.Raw.Commit.ID)
		}

		fmt.Printf("Main hash: %s\n", mainHash)
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
				fmt.Printf(
					"Skip pipeline (sematic): %s %s\n",
					i.Ref,
					i.SHA,
				)
			}

			continue
		}

		if latestMain != nil &&
			i.Ref == mainBranch.Name &&
			i.SHA == latestMain.SHA {
			if o.Verbose {
				fmt.Printf(
					"Skip pipeline (main): %s %s\n",
					i.Ref,
					i.SHA,
				)
			}

			continue
		}

		if o.Verbose {
			fmt.Printf("Pipeline: %s %s\n", i.Ref, i.SHA)
		} else {
			fmt.Printf("Pipeline: %s\n", i.Ref)
		}

		c.MustDeletePipeline(p.Identifier, i.ID)
	}
}
