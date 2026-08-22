package transcript_cache

func (c *Cache) Delete(sessionIdentifier string) {
	c.mutex.Lock()
	delete(c.sessions, sessionIdentifier)
	delete(c.calls, sessionIdentifier)
	c.mutex.Unlock()
	c.Client.Delete(sessionIdentifier)
}
