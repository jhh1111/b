#include <iostream>
#include <stdlib.h>
#include <string.h>
#include <string>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <event.h>
#include <jsoncpp/json/json.h>
#include<mysql/mysql.h>

using namespace std;
const int lis_max=20;
enum OP_TYPE
{
    DL = 1,
    ZC,
    CKYY,
    YD,
    YDXX,
    QXYD,
    TC
};

class MyLibevent;
class Socket
{
public:
    Socket()
    {
        m_ips = "127.0.0.1";
        m_port = 6000;
        m_sockfd = -1;
    }
    Socket(const string &ips, short port) : m_ips(ips), m_port(port)
    {
        m_sockfd = -1;
    }
    bool Socket_init();
    int Get_sockfd() const
    {
        return m_sockfd;
    }
    ~Socket()
    {
        close(m_sockfd);
    }

private:
    string m_ips;
    short m_port;
    int m_sockfd;
};

class Sock_Obj
{
public:
    virtual void CallBack_Fun()=0;
    struct event* ev;
    MyLibevent* plib;
};
class MysqlClient
{
public:
    MysqlClient()
    {
        ips="127.0.0.1";
        mysql_username="root";
        mysql_userpasswd="123456";
        mysql_dbname="c2301";
        port=3306;
    }
    bool Connect_toDb();
    bool Db_user_zc(string name,string tel,string pw);
    bool Db_user_dl(string &name,string tel,string pw);
    bool Db_show_yuyue(Json::Value &resval);
    bool Db_user_yd(string user_tel,string tk_id );
    bool Db_show_yd(Json::Value& resval,string user_tel);
    bool Db_user_delyd(string res_id);

    void Begin();
    void RollBack();
    void Commit();

    ~MysqlClient()
    {
        mysql_close(&mysql_con);
    }
private:
    MYSQL mysql_con;
    string ips;
    string mysql_username;
    string mysql_userpasswd;
    string mysql_dbname;
    short port;
};
class Accept_Obj:public Sock_Obj
{
public:
    Accept_Obj(int fd, MyLibevent* p)
    {
        sockfd=fd;
        plib=p;
    }
    virtual void CallBack_Fun();
private:
    int sockfd;
};
class Recv_Obj:public Sock_Obj
{
public:
    Recv_Obj(int fd,MyLibevent* p):c(fd)
    {
        plib=p;
    }
    virtual void CallBack_Fun();
    ~Recv_Obj()
    {
        close(c);
        cout<<"client close"<<endl;
    }
private:
    void Send_OK();
    void Send_ERR();
    void Send_Json(Json::Value & val);

    void User_zc();
    void User_dl();

    void show_YuYue();
    void User_yd(); 
    void Show_user_yd();
    void Delete_user_yd();
private:
    int c;
    Json::Value m_val;

};
class MyLibevent
{
public:
    MyLibevent()
    {
        m_base=NULL;
    }
    bool MyLibevent_Init()
    {
        m_base=event_init();
        if(m_base==NULL)
        {
            return false;
        }
        return true;
    }
    bool MyLibevent_Add(int fd,Sock_Obj* pObj);
    void MyLibevent_Delete(Sock_Obj* pObj);
    void MyLibevent_Loop()
    {
        if(m_base!=NULL)
        {
            event_base_dispatch(m_base);
        }
    }

private:
    struct event_base* m_base;
};