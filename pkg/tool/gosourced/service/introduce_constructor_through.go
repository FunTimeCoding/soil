package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern_site"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"os"
	"path/filepath"
)

func (s *Service) introduceConstructorThrough(
	directory string,
	packagePath string,
	name string,
	parameters []string,
	fromSites bool,
	assignRest bool,
	out *sink.Sink,
) (*output.Results, *result.Constructor, error) {
	r := output.NewResultsWithDirectory(directory)
	refuse := func(message string) (*output.Results, *result.Constructor, error) {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, message, "", false),
		)

		return r, nil, nil
	}

	if fromSites && len(parameters) > 0 {
		return refuse("parameters and parameters_from_sites exclude each other")
	}

	all, set, e := s.censusPackages(directory, packagePath)

	if e != nil {
		return nil, nil, e
	}

	declaration, p, f := findDeclaration(all, packagePath, name, "")

	if f != nil {
		return refuse(f.Error())
	}

	named, structure := namedStruct(declaration)

	if structure == nil {
		return refuse(
			fmt.Sprintf("%s is not a struct type in %s", name, packagePath),
		)
	}

	if named.TypeParams().Len() > 0 {
		return refuse(
			fmt.Sprintf("%s is generic - write its constructor by hand", name),
		)
	}

	constructor, file, refusal := constructorHome(p, name)

	if refusal != "" {
		return refuse(refusal)
	}

	census := literalCensus(all, set, named, structure, packagePath)

	if fromSites {
		parameters = siteIntersection(census.Sites)
	}

	chosen, g := constructorParameters(structure, parameters)

	if g != nil {
		return refuse(g.Error())
	}

	source, h := constructorSource(p.Types, name, constructor, chosen)

	if h != nil {
		return refuse(h.Error())
	}

	path := filepath.Join(filepath.Dir(p.GoFiles[0]), file)

	if _, i := os.Stat(path); i == nil {
		return refuse(fmt.Sprintf("%s already exists", path))
	}

	var names []string

	for _, v := range chosen {
		names = append(names, v.Name())
	}

	report := result.NewConstructor(
		name,
		constructor,
		system.RelativePath(directory, path),
		names,
	)
	report.Total = census.Total
	report.Inside = census.Inside
	report.Expected = census.Expected
	decorations := decoration.NewSet()
	contents := map[string][]byte{}
	var remaining []*pattern_site.Entry

	for _, site := range census.Sites {
		reason, j := rewriteSite(
			decorations,
			set,
			structure,
			site,
			chosen,
			packagePath,
			constructor,
			assignRest,
		)

		if j != nil {
			return nil, nil, j
		}

		if reason == "" {
			report.Rewritten++

			continue
		}

		entry, k := literalEntry(
			directory,
			set,
			contents,
			site,
			fmt.Sprintf("%s - %s", site.Shape, reason),
		)

		if k != nil {
			return nil, nil, k
		}

		remaining = append(remaining, entry)
	}

	report.Remaining = groupEntries(remaining)
	out.Write(path, source)
	resolver := resolve.NewNames(all)

	for filename, decorated := range decorations.Files {
		if m := restoreDecoratedFile(
			resolver,
			decorations.PackagePaths[decorated],
			map[string]string{},
			decorated,
			filename,
			out,
		); m != nil {
			return nil, nil, m
		}
	}

	return r, report, nil
}
