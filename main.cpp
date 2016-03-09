using namespace std;

#include<iostream>
#include "lib\ArrayList.cpp"
#include "util\ListUtil.cpp"


int main()
{
	ArrayList a(6);	
	ListUtil::readList(&a);
	ListUtil::printList(&a);
	
	return 0;
}
