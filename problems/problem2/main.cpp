/*
 * Author: Duy Quoc
 *
 * Description: Parse characters from file with utf-8 encoding
 *
 * Detail: 
 *      - English
 *      - Vietnamese
 *      - Chinese
 */

using namespace std;

#include<iostream>
#include<cstdint>

unsigned char* text = new unsigned char[7];

void readString()
{
    //sample text: đặng
    //code point: đ     U+0111  - 273   - 0001 0001 0001
    //  code unit       0xC4    - 196   - 11000100 
    //                  0x91    - 145   - 10010001
    //
    //code point: ặ     U+1EB7  - 7863  - 0001 1110 1011 0111
    //  code unit       0xE1    - 225   - 1110 0001
    //                  0xBA    - 186   - 1011 1010
    //                  0xB7    - 183   - 1011 0111
    //
    //code point: n     U+006E  - 110
    //  code unit       0x6E    - 110   - 0110 1110
    //
    //code point: g     U+0067  - 103
    //  code unit       0x67    - 103   - 0110 0111

    text[0] = 0xC4;
    text[1] = 0x91;
    text[2] = 0xE1;
    text[3] = 0xBA;
    text[4] = 0xB7;
    text[5] = 0x6E;
    text[6] = 0x67;

}

void parseString() 
{
}

void printString()
{
    cout<< "This string is:" << endl;
    for(int index = 0; index < 7; ++index)
    {
        cout<< (int)text[index] << ", ";
    }

    cout<< endl;

    cout<< 0x0111 << " - " << 0x1EB7 << "- " << 0x006E << " - " << 0x0067 << endl;
}


int main()
{

    unsigned char mask1 = 224;          //1110 0000
    unsigned char mask2 = 31;           //0001 1111

    unsigned char pattern1 = 223;       //1101 1111      
    unsigned char pattern2 = 143;       //1000 1111

    uint32_t p = pattern1 & mask1; 

    cout<< (p << 1) << "" << (pattern1 & mask2) << endl;

    readString();
    printString();

    return 0;
}
