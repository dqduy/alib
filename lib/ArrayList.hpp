
class ArrayList
{
	private:
		int maxLength; //Max item of the list
		int cursor; //Current cursor 
		
	public: 
		ArrayList(int maxLength);
		bool is Empty();
		void clear();
		void add(int location, int item);
		void add(int item);
		void remove(int location);
		int indexOf(int item);
		int get(int location);
		int size();		
};



