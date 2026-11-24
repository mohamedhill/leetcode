func isIsomorphic(s string, t string) bool {
    mapa1 := map[byte]byte{}
    mapa2 := map[byte]byte{}

  

    for i := 0; i < len(s); i++ {
       

        mapa1[s[i]] = t[i]
        mapa2[t[i]]=s[i]
       
    }
    fmt.Println(mapa1)
    for i := 0 ; i <len(t);i++{
        if t[i]!=mapa1[s[i]]{

            return false 
        }

        if s[i]!=mapa2[t[i]]{
            return false
        }


    }



    return true
}