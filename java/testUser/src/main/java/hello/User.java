package hello;

import javax.persistence.Entity;

import javax.persistence.Table;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

@Entity
@Table(name = "testtable")
public class User {

	@Id
	@GeneratedValue(strategy = GenerationType.AUTO)
	private long idtesttable;
	
    private String maketxt;
	
	public User(){
		
	}

    public User(long idtesttable, String maketxt) {
        this.idtesttable = idtesttable;
		this.maketxt = maketxt;
    }

    public long getIdtesttable() {
        return idtesttable;
    }
	
	public void setIdtesttable(long idtesttable){
		this.idtesttable = idtesttable;
	}

    public String getMaketxt() {
        return maketxt;
    }
	public void setContent(String maketxt){
		this.maketxt = maketxt;
	}
}
