package readme

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

var (
	ErrLinkKeyInUse         = errors.New("link key is already in use")
	ErrToDeepSectionNesting = errors.New("too deep nesting")
)

type File struct {
	Path   string
	Config Config

	lines []string
	links map[string]Link
}

func New() *File {
	r := &File{}
	r.links = make(map[string]Link)
	return r
}

func (r *File) ReadConfigBytes(config []byte) error {
	if err := yaml.Unmarshal(config, &r.Config); err != nil {
		return err
	}
	if len(r.Config.Links) > 0 {
		for key, link := range r.Config.Links {
			if err := r.AddLink(key, link.Link, link.Title); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *File) AddLink(key, link, title string) error {
	if _, ok := r.links[key]; ok {
		return fmt.Errorf("key (%s): %w", key, ErrLinkKeyInUse)
	}
	r.links[key] = Link{key, link, title}
	return nil
}

func (r *File) Render() (content []byte, err error) {
	// readme file comment
	comment := []string{
		"<!--",
		"DO NOT EDIT. This README is generated by readme generator.",
		"Generated using cmd './github-cicd-experiments readme update'",
		"and config file which can be found here",
		"./cmd/github-cicd-experiments/internal/assets/readme.yml",
		fmt.Sprintf("Generated at: %s", time.Now().UTC()),
		"-->",
		"",
	}
	r.lines = append(comment, r.lines...)

	// intro
	intro := []string{}
	if len(r.Config.Title) > 0 {
		intro = append(intro, "# "+r.Config.Title+"\n")
	}
	if len(r.Config.Markdown) > 0 {
		intro = append(intro, r.Config.Markdown)
	}
	r.lines = append(r.lines, intro...)

	// toc
	toc := []string{
		"# Table of Contents",
		"",
	}

	// content
	var sections []string
	for _, section := range r.Config.Content {
		// start from h2 (##)
		slines, stoc, err := section.Lines(1)
		if err != nil {
			return content, err
		}
		sections = append(sections, slines...)

		if len(stoc) > 0 {
			toc = append(toc, stoc...)
		}
	}

	r.lines = append(r.lines, toc...)
	r.lines = append(r.lines, "")

	r.lines = append(r.lines, sections...)

	// links
	if len(r.links) > 0 {
		r.lines = append(r.lines, "<!-- LINKS -->")
		for _, link := range r.links {
			r.lines = append(r.lines, link.Line())
		}
	}
	r.lines = append(r.lines, "")
	contentStr := strings.Join(r.lines, "\n")
	content = []byte(contentStr)
	return content, err
}

type Config struct {
	Title    string           `yaml:"title"`
	Markdown string           `yaml:"markdown"`
	Content  []ContentSection `yaml:"content"`
	Links    map[string]Link  `yaml:"links"`
}

type ContentSection struct {
	Title    string           `yaml:"title"`
	Markdown string           `yaml:"markdown"`
	Table    ContentTable     `yaml:"table"`
	Sections []ContentSection `yaml:"sections"`
}

func (cs ContentSection) Lines(hdepth int) (lines, toc []string, err error) {
	if hdepth > 10 {
		return nil, nil, ErrToDeepSectionNesting
	}
	if len(cs.Title) > 0 {
		if hdepth <= 6 {
			lines = append(lines, fmt.Sprintf("%s %s\n", strings.Repeat("#", hdepth), cs.Title))
			toc = append(toc, fmt.Sprintf("%s- [%s](#%s)", strings.Repeat("  ", hdepth-1), cs.Title, SlugOf(cs.Title)))
		} else {
			lines = append(lines, fmt.Sprintf("**%s**\n", cs.Title))
		}
	}

	if len(cs.Markdown) > 0 {
		lines = append(lines, cs.Markdown)
	}

	lines = append(lines, cs.Table.Lines()...)

	next := hdepth + 1
	for _, section := range cs.Sections {
		slines, stoc, err := section.Lines(next)
		if err != nil {
			return nil, nil, err
		}
		lines = append(lines, slines...)
		// only add sub section toc if parent section has toc def
		if len(toc) > 0 {
			toc = append(toc, stoc...)
		}

	}

	if len(lines) > 0 {
		if last := lines[len(lines)-1]; last != "" {
			lines = append(lines, "")
		}
	}

	return
}

type ContentTable struct {
	Cols []string          `yaml:"cols"`
	Rows []ContentTableRow `yaml:"rows"`
}

type ContentTableRow struct {
	Cols []string `yaml:"cols"`
}

func (table ContentTable) Lines() (lines []string) {
	if len(table.Cols) == 0 {
		return
	}
	lines = append(lines, "| "+strings.Join(table.Cols, " | ")+" |")
	lines = append(lines, strings.Repeat("| --- ", len(table.Cols))+"|")

	for _, row := range table.Rows {
		lines = append(lines, "| "+strings.Join(row.Cols, " | ")+" |")
	}
	return
}

type Link struct {
	Key   string `yaml:"key"`
	Link  string `yaml:"link"`
	Title string `yaml:"title"`
}

func (l Link) Line() string {
	link := fmt.Sprintf("[%s]: %s", l.Key, l.Link)
	if len(l.Title) > 0 {
		link += fmt.Sprintf(" %q", l.Title)
	}
	return link
}
