package model


import "demo/config/db"

// import "fmt"

type Site struct {
	Site_key string 
	Id int

}

func AllSites() []*Site{
	var err error
	rows := db.ExecSql("select site_key, id from sites ")
	var sites []*Site
	defer rows.Close()
	for rows.Next() {
		site := new(Site)
		err  = rows.Scan(&site.Site_key, &site.Id)
		raiseError(err)
		sites = append(sites, site)
	}
	return sites
}

func raiseError(err error){
	if err != nil{
		panic(err)
	}
}