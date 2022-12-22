#ifndef SERIALIZATION_H
#define SERIALIZATION_H

#include <iostream>
#include <thrift/TApplicationException.h>
#include <thrift/protocol/TCompactProtocol.h>
#include <thrift/protocol/TBinaryProtocol.h>
#include <thrift/transport/TBufferTransports.h>
#include <vector>

inline std::shared_ptr<apache::thrift::protocol::TProtocol> create_deserialize_protocol(
        std::shared_ptr<apache::thrift::transport::TMemoryBuffer> mem, 
        bool compact)
{
    if (compact) {
        apache::thrift::protocol::TCompactProtocolFactoryT<apache::thrift::transport::TMemoryBuffer>
                tproto_factory;
        return tproto_factory.getProtocol(mem);
    } else {
        apache::thrift::protocol::TBinaryProtocolFactoryT<apache::thrift::transport::TMemoryBuffer>
                tproto_factory;
        return tproto_factory.getProtocol(mem);
    }
}

template<typename ThriftT> 
ThriftT deserializeFromBinanry(void* buffer, uint32_t size)
{
    uint8_t* binary_buffer = static_cast<uint8_t*>(buffer);
    std::shared_ptr<apache::thrift::transport::TMemoryBuffer> tmem_transport(
            new apache::thrift::transport::TMemoryBuffer(binary_buffer, size));
    std::shared_ptr<apache::thrift::protocol::TProtocol> proto = create_deserialize_protocol(tmem_transport, false);

    ThriftT thrift_type;
    try
    {
        thrift_type.read(proto.get());
    } 
    catch (std::exception& e)
    {
        std::cout << "Couldn't deserialize thrift msg: " << e.what() << std::endl;
        exit(-1);
    }

    return thrift_type;
}


class ThriftSerializer {
public:
    // If compact, the objects will be serialized using the Compact Protocol.  Otherwise,
    // we'll use the binary protocol.
    // Note: the deserializer must be matching.
    ThriftSerializer(bool compact, int initial_buffer_size) :
        mem_buffer(new apache::thrift::transport::TMemoryBuffer(initial_buffer_size)) {
        if (compact) {
            apache::thrift::protocol::TCompactProtocolFactoryT<apache::thrift::transport::TMemoryBuffer>
                    factory;
            protocol = factory.getProtocol(mem_buffer);
        } else {
            apache::thrift::protocol::TBinaryProtocolFactoryT<apache::thrift::transport::TMemoryBuffer>
                    factory;
            protocol = factory.getProtocol(mem_buffer);
        }
    }

    // Serializes obj into result. Result will contain a copy of the memory.
    template <class T>
    bool serialize(T* obj, std::vector<uint8_t>* result) {
        uint32_t len = 0;
        uint8_t* buffer = nullptr;
        RETURN_IF_ERROR(serialize<T>(obj, &len, &buffer));
        result->resize(len);
        memcpy(&((*result)[0]), buffer, len);
        return true;
    }

    // serialize obj into a memory buffer. The result is returned in buffer/len. The
    // memory returned is owned by this object and will be invalid when another object
    // is serialized.
    template <class T>
    bool serialize(T* obj, uint32_t* len, uint8_t** buffer) {
        try {
            mem_buffer->resetBuffer();
            obj->write(protocol.get());
        } catch (std::exception& e) {
            std::cout << "Couldn't serialize thrift object:\n" << e.what();
            return false;
        }

        mem_buffer->getBuffer(buffer, len);
        return true;
    }

    template <class T>
    bool serialize(T* obj, std::string* result) {
        try {
            mem_buffer->resetBuffer();
            obj->write(protocol.get());
        } catch (apache::thrift::TApplicationException& e) {
            std::cout << "Couldn't serialize thrift object:\n" << e.what();
            return false;
        }

        *result = mem_buffer->getBufferAsString();
        return true;
    }

    template <class T>
    bool serialize(T* obj) {
        try {
            mem_buffer->resetBuffer();
            obj->write(protocol.get());
        } catch (apache::thrift::TApplicationException& e) {
            std::cout << "Couldn't serialize thrift object:\n" << e.what();
            return false;
        }

        return true;
    }

    void get_buffer(uint8_t** buffer, uint32_t* length) { mem_buffer->getBuffer(buffer, length); }

private:
    std::shared_ptr<apache::thrift::transport::TMemoryBuffer> mem_buffer;
    std::shared_ptr<apache::thrift::protocol::TProtocol> protocol;
};

#endif