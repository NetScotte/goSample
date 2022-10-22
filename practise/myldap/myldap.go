package myldap

import "gopkg.in/ldap.v2"

type LdapClient struct {
	Server string
	Dn     string
	Passwd string
	BaseDn string
	Conn   *ldap.Conn
}

func GetLdapClient(server, dn, passwd, baseDn string) (*LdapClient, error) {
	c, err := ldap.Dial("tcp", server)
	if err != nil {
		return nil, err
	}
	err = c.Bind(dn, passwd)
	if err != nil {
		return nil, err
	}
	ldapClient := &LdapClient{
		Server: server,
		Dn:     dn,
		Passwd: passwd,
		BaseDn: baseDn,
		Conn:   c,
	}
	return ldapClient, nil
}

func (l *LdapClient) SearchByFilter(filter string) (*ldap.SearchResult, error) {
	// filter: (&(objectClass=organizationalPerson)(uid=sco))
	searchRequest := ldap.NewSearchRequest(l.BaseDn, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter, []string{"dn"}, nil)
	return l.Conn.Search(searchRequest)
}
