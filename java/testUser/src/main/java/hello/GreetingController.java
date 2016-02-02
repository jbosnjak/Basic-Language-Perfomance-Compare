package hello;

import java.util.concurrent.atomic.AtomicLong;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.beans.factory.annotation.Autowired;

@RestController
public class GreetingController {

    private final AtomicLong counter = new AtomicLong();

    @RequestMapping("/listUsers")
    public String listUsers(@RequestParam(value="user", defaultValue="1") long id) {
        
		try {
			User user = userDao.findOne(id);
			if (user == null){
				return "User Not Found";
			}
			return user.toString();
		}
		catch (Exception ex) {
			ex.printStackTrace();
			return "User Not Found";
		}
    }
	
	 @Autowired
	private UserDao userDao;
}
