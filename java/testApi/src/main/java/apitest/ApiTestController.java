package apitest;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.beans.factory.annotation.Autowired;

@RestController
public class ApiTestController {

    @RequestMapping("/testApiDB")
    public String testApiDB(@RequestParam(value="user", defaultValue="1") long id) {
        
		try {
			User user = userDao.findOne(id);
			if (user == null){
				return "User Not Found";
			}
			return user.getContent().toString();
		}
		catch (Exception ex) {
			ex.printStackTrace();
			return "User Not Found";
		}
    }
    
    @RequestMapping("/testApiPing")
    public String testApiPing(@RequestParam(value="param1", defaultValue="1") String id) {
        
		return new StringBuilder(id).reverse().toString();
    }
	
	 @Autowired
	private UserDao userDao;
}
