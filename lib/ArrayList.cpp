#include "ArrayList.h"

ArrayList::ArrayList(int maxLength) 
{
	list = new int[maxLength];
	this->maxLength = maxLength;
	this->cursor = 0;
}

bool ArrayList::isEmpty()
{
	if(cursor == 0) return true;
	else return false;
}

bool ArrayList::isFull()
{
	if(cursor == maxLength) return true;
	else return false;
}

void ArrayList::clear()
{
	delete[] list;
	cursor = 0;
}

void ArrayList::add(int location, int item)
{
	if(!isFull()) 
	{
		if(location >= 1 && location <= cursor + 1)
		{
			for(int i = cursor; i> location-1; i--)
			{
				list[i] = list[i-1];
			}
			
			list[location-1] = item;
			cursor++;
		}
	}
}

void ArrayList::add(int item)
{
	if(!isFull())
	{
		list[cursor] = item;
		cursor++;
	}
}

void ArrayList::remove(int location)
{
	
} 
 
int ArrayList::indexOf(int item)
{
	return 0;
} 
 
int ArrayList::get(int location)
{
	return list[location-1];
} 

int ArrayList::size()
{
	return cursor;
}
