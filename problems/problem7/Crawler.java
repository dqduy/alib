import java.net.*;
import java.net.*;
import java.io.*;
import java.util.*;

public class Crawler {
	public int totalItem = 0;
	public int chunk = 10000;
	
	public void init() throws Exception{
		this.totalItem = getTotalItem();
	}
	
	public int getTotalItem() throws Exception{
		URL oracle = new URL("https://hacker-news.firebaseio.com/v0/maxitem.json?print=pretty");
		BufferedReader in = new BufferedReader(new InputStreamReader(oracle.openStream()));
		String inputLine = "";
		StringBuilder builder = new StringBuilder();
		int totalItem = 0;

		while ((inputLine = in.readLine()) != null) {
			//builder.append(inputLine);
			totalItem = Integer.parseInt(inputLine);
		}
				
		in.close();
		
		//System.out.println("totalItem: " + totalItem + " - " + builder.toString());
		return totalItem;
	}
	
	public String getItem(int id) throws Exception{
		URL oracle = new URL("https://hacker-news.firebaseio.com/v0/item/" + id + ".json?print=pretty");
		BufferedReader in = new BufferedReader(new InputStreamReader(oracle.openStream()));
		String inputLine = "";
		StringBuilder builder = new StringBuilder();

		while ((inputLine = in.readLine()) != null) {
			System.out.println(inputLine);
			builder.append(inputLine);
		}
				
		in.close();
		
		return builder.toString();
	}
	
	public void writeToDisk(String content, String path) {
		File file = new File(path);
		
		try (FileOutputStream fop = new FileOutputStream(file)) {

			// if file doesn't exists, then create it
			if (!file.exists()) {
				file.createNewFile();
			}

			// get the content in bytes
			byte[] contentInBytes = content.getBytes();

			fop.write(contentInBytes);
			fop.flush();
			fop.close();

			System.out.println("Done");

		} catch (IOException e) {
			e.printStackTrace();
		}		
	}
	
	public void craw() throws Exception{
		StringBuilder b = new StringBuilder();
		String s = "";
		long startTick = System.currentTimeMillis();
		
		for(int index = 1; index <= 1000; ++index) {
			s = getItem(index);
			//System.out.println();
			/* if(index % chunk == 0)
				writeToDisk(s, index + ".txt");
			if(index == totalItem)
				writeToDisk(s, index + ".txt"); */
		}
		
		System.out.println("Time eslapse: " + (System.currentTimeMillis() - startTick));		
	}
	
	public void crawMulti() throws InterruptedException{
		long startTick = System.currentTimeMillis();
		List<Thread> threads = new ArrayList<>();
		
		for(int index = 0; index < 10; ++index) {
			CrawThread t = new CrawThread("" + index);
			t.start();
			threads.add(t);			
		}
		
		for (Thread t : threads) {
			t.join();
		}
		
		System.out.println("Time eslapse with multithread: " + (System.currentTimeMillis() - startTick));		
	}
	
    public static void main(String[] args) throws Exception{
		System.getProperties().put("http.proxyHost", "");
		System.getProperties().put("http.proxyPort", "");
		//System.getProperties().put("http.proxyUser", "");
		//System.getProperties().put("http.proxyPassword", "");
		Crawler c = new Crawler();
		
		// c.init();
		//c.craw();
		c.crawMulti();
    }
}

 class CrawThread extends Thread {
	private String id;
	 
	public CrawThread() {
	}
	
	public CrawThread(String id) {
		this.id = id;
	}

	public String getItem(int id) throws Exception{
		URL oracle = new URL("https://hacker-news.firebaseio.com/v0/item/" + id + ".json?print=pretty");
		BufferedReader in = new BufferedReader(new InputStreamReader(oracle.openStream()));
		String inputLine = "";
		StringBuilder builder = new StringBuilder();

		while ((inputLine = in.readLine()) != null) {
			System.out.println(inputLine);
			builder.append(inputLine);
		}
				
		in.close();
		
		return builder.toString();
	}
	
	public void craw() throws Exception {
		String s = "";
		for(int index = 1; index <= 1000; ++index) {
			s = getItem(index);
		}		
	}
	
	public void run(){
		try {
			craw();
		}
		catch(Exception e) {
			
		}
		
	}
}