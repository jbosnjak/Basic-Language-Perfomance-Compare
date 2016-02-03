package apitest;

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
	
    private String content;
	
	public User(){
		
	}

    public User(long idtesttable, String content) {
        this.idtesttable = idtesttable;
		this.content = content;
    }

    public long getIdtesttable() {
        return idtesttable;
    }
	
	public void setIdtesttable(long idtesttable){
		this.idtesttable = idtesttable;
	}

    public String getContent() {
        return content;
    }
	public void setContent(String content){
		this.content = content;
	}
}
