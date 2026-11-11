package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/resolve_result"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/match"
)

func (s *Store) ResolveSessionIdentifier(query string) (*resolve_result.Result, error) {
	seen := map[string]bool{}
	var matches []*match.Match
	var callsignSession session.Session
	r := s.database.Where("callsign = ?", query).Limit(1).Find(&callsignSession)

	if r.Error != nil {
		return nil, r.Error
	}

	if r.RowsAffected > 0 {
		seen[callsignSession.Identifier] = true
		matches = append(
			matches,
			matchFrom(&callsignSession, constant.Callsign),
		)
	}

	var identifierSession session.Session
	r = s.database.Where("identifier = ?", query).Limit(1).Find(
		&identifierSession,
	)

	if r.Error != nil {
		return nil, r.Error
	}

	if r.RowsAffected > 0 && !seen[identifierSession.Identifier] {
		seen[identifierSession.Identifier] = true
		matches = append(
			matches,
			matchFrom(&identifierSession, constant.Identifier),
		)
	}

	var aliasSession session.Session
	r = s.database.Where("alias = ?", query).Limit(1).Find(&aliasSession)

	if r.Error != nil {
		return nil, r.Error
	}

	if r.RowsAffected > 0 && !seen[aliasSession.Identifier] {
		seen[aliasSession.Identifier] = true
		matches = append(matches, matchFrom(&aliasSession, constant.Alias))
	}

	var slugSessions []session.Session

	if e := s.database.Where("slug = ?", query).Find(&slugSessions).Error; e != nil {
		return nil, e
	}

	for i := range slugSessions {
		if !seen[slugSessions[i].Identifier] {
			seen[slugSessions[i].Identifier] = true
			matches = append(
				matches,
				matchFrom(&slugSessions[i], constant.Slug),
			)
		}
	}

	var prefixSessions []session.Session

	if e := s.database.Where(
		"identifier LIKE ?",
		fmt.Sprintf("%s%%", query),
	).Find(
		&prefixSessions,
	).Error; e != nil {
		return nil, e
	}

	for i := range prefixSessions {
		if !seen[prefixSessions[i].Identifier] {
			seen[prefixSessions[i].Identifier] = true
			matches = append(matches, matchFrom(&prefixSessions[i], "prefix"))
		}
	}

	return resolve_result.New(matches), nil
}
