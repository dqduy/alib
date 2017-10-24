#ifndef __CODEPOINT_H_
#define __CODEPOINT_H_

#include<vector>
#include<cstdint>

class U8Mask {
    public:
        //Bit mask to filter data from byte
        static const uint8_t maskClass1e = 0x80;     //1000 0000
        static const uint8_t maskClass2e = 0xE0;     //1110 0000
        static const uint8_t maskClass3e = 0xF0;     //1111 0000
        static const uint8_t maskClass4e = 0xF8;     //1111 1000
        static const uint8_t maskClassee = 0xC0;     //1100 0000

        static const uint8_t maskClass1d = 0x7F;     //0111 1111
        static const uint8_t maskClass2d = 0x1F;     //0001 1111
        static const uint8_t maskClass3d = 0x0F;     //0000 1111
        static const uint8_t maskClass4d = 0x07;     //0000 0111
        static const uint8_t maskClassed = 0x3F;     //0011 1111
        //End bit mask
};

class CodePoint {
    private:
        std::vector<uint8_t> codeUnits;
    
    public:
        CodePoint();
        void addCodeUnit(uint8_t unit);
        uint32_t getCodePoint();
        int size();
        uint8_t* getContent();
        uint8_t* getString();
        void print();
};

#endif


