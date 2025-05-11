#include "ser.h"

void G_CALLBACK_FUN(int fd, short ev, void *arg)
{
    Sock_Obj *ptr = (Sock_Obj *)arg;
    if (ptr == NULL)
    {
        return;
    }
    if (ev & EV_READ)
    {
        ptr->CallBack_Fun();
    }
}
bool Socket::Socket_init()
{
    m_sockfd = socket(AF_INET, SOCK_STREAM, 0);
    if (m_sockfd == -1)
    {
        return false;
    }

    struct sockaddr_in saddr;
    memset(&saddr, 0, sizeof(saddr));
    saddr.sin_family = AF_INET;
    saddr.sin_port = htons(m_port);
    saddr.sin_addr.s_addr = inet_addr(m_ips.c_str());
    if (bind(m_sockfd, (struct sockaddr *)&saddr, sizeof(saddr)) == -1)
    {
        return false;
    }
    if (listen(m_sockfd, lis_max) == 1)
    {
        return false;
    }
    return true;
}
bool MyLibevent::MyLibevent_Add(int fd, Sock_Obj *pObj)
{
    if (pObj == NULL)
    {
        return false;
    }
    struct event *pev = event_new(m_base, fd, EV_READ | EV_PERSIST, G_CALLBACK_FUN, pObj);
    if (pev == NULL)
    {
        return false;
    }
    pObj->ev = pev;
    event_add(pev, NULL);
    return true;
}
void MyLibevent::MyLibevent_Delete(Sock_Obj *pObj)
{
    if (pObj != NULL)
    {
        event_free(pObj->ev);
    }
}
void MysqlClient::Begin()
{
    if(mysql_query(&mysql_con,"begin")!=0)
    {
        cout<<"开始事务失败"<<endl;
    }
}
void MysqlClient::RollBack()
{
    if(mysql_query(&mysql_con,"rollback")!=0)
    {
        cout<<"回滚失败"<<endl;
    }
}
void MysqlClient::Commit()
{
    if(mysql_query(&mysql_con,"commit")!=0)
    {
        cout<<"提交事务失败"<<endl;
    }
}
bool MysqlClient::Connect_toDb()
{
    MYSQL *mysql = mysql_init(&mysql_con);
    if (mysql == NULL)
    {
        return false;
    }
    mysql = mysql_real_connect(mysql, ips.c_str(), mysql_username.c_str(), mysql_userpasswd.c_str(), mysql_dbname.c_str(), port, NULL, 0);
    if (mysql == NULL)
    {
        cout << "连接数据库失败" << endl;
        return false;
    }
    return true;
}
bool MysqlClient::Db_user_zc(string name, string tel, string pw)
{
    string sql = string("insert into user_info value(0,'") + tel + string("','") + name + string("','") + pw + string("','1')");
    if (mysql_query(&mysql_con, sql.c_str()) != 0)
    {
        return false;
    }
    return true;
}
bool MysqlClient::Db_user_dl(string &name, string tel, string pw)
{
    string sql = string("select user_name,user_passwd from user_info where user_tel=") + tel;

    if (mysql_query(&mysql_con, sql.c_str()) != 0)
    {
        return false;
    }
    MYSQL_RES *r = mysql_store_result(&mysql_con);
    if (r == NULL)
    {
        return false;
    }
    int num = mysql_num_rows(r);
    if (num != 1)
    {
        return false;
    }
    MYSQL_ROW row = mysql_fetch_row(r);
    name = row[0];
    string passwd = row[1];
    if (passwd.compare(pw) != 0)
    {
        return false;
    }
    return true;
}
bool MysqlClient::Db_show_yuyue(Json::Value &resval)
{
    string sql = string("select * from ticket_info");
    if (mysql_query(&mysql_con, sql.c_str()) != 0)
    {
        return false;
    }
    MYSQL_RES *r = mysql_store_result(&mysql_con);
    if (r == NULL)
    {
        return false;
    }
    int num = mysql_num_rows(r);
    resval["num"] = num;
    if (num == 0)
    {
        return true;
    }
    for (int i = 0; i < num; i++)
    {
        MYSQL_ROW row = mysql_fetch_row(r);
        Json::Value tmp;
        tmp["tk_id"] = row[0];
        tmp["tk_name"] = row[1];
        tmp["tk_max"] = row[2];
        tmp["tk_count"] = row[3];
        tmp["tk_date"] = row[4];

        resval["arr"].append(tmp);
    }
    return true;
}
bool MysqlClient::Db_user_yd(string user_tel, string tk_id)
{
    // select tk_max,tk_count from ticket_info where tk_id=1
    string sql_read = string("select tk_max,tk_count from ticket_info where tk_id=") + tk_id;
    if (mysql_query(&mysql_con, sql_read.c_str()) != 0)
    {
        return false;
    }
    MYSQL_RES *r = mysql_store_result(&mysql_con);
    if (r == NULL)
    {
        return false;
    }
    int num = mysql_num_rows(r);
    if (num != 1)
    {
        return false;
    }
    MYSQL_ROW row = mysql_fetch_row(r);
    // row[0] row[1]
    int max_num = atoi(row[0]);
    int count = atoi(row[1]);
    if (count >= max_num)
    {
        return false;
    }
    count++;
    Begin();
    string sql_update = string("update ticket_info set tk_count=") + to_string(count) + string(" ") + string("where tk_id=") + tk_id;
    if (mysql_query(&mysql_con, sql_update.c_str()) != 0)
    {
        RollBack();
        return false;
    }
    // insert into ticket_res values(0,1,13500000000,now());

    string sql_res = string("insert into ticket_res values(0,") + tk_id + string(",") + user_tel + string(",now())");
    if (mysql_query(&mysql_con, sql_res.c_str()) != 0)
    {
        RollBack();
        return false;
    }
    Commit();
    return true;
}
bool MysqlClient::Db_show_yd(Json::Value &resval, string user_tel)
{
    string sql = string("select res_id,tk_name,yd_time from ticket_info,ticket_res where ticket_info.tk_id = ticket_res.tk_id and ticket_res.user_tel=") + user_tel;
    if (mysql_query(&mysql_con, sql.c_str()) != 0)
    {
        return false;
    }
    MYSQL_RES *r = mysql_store_result(&mysql_con);
    if (r == NULL)
    {
        return false;
    }
    int num = mysql_num_rows(r);
    resval["num"] = num;
    if (num == 0)
    {
        return true;
    }
    for (int i = 0; i < num; i++)
    {
        MYSQL_ROW row = mysql_fetch_row(r);
        Json::Value tmp;
        tmp["res_id"] = row[0];
        tmp["tk_name"] = row[1];
        tmp["yd_time"] = row[2];

        resval["arr"].append(tmp);
    }
    return true;
}
bool MysqlClient::Db_user_delyd(string res_id)
{
    string sql_tkid=string("select tk_id from ticket_res where res_id=")+res_id;
    if(mysql_query(&mysql_con,sql_tkid.c_str())!=0)
    {
        cout<<"查询tk_id失败"<<endl;
        return false;
    }
    MYSQL_RES *r =mysql_store_result(&mysql_con);
    if(r==NULL)
    {
        return false;
    }
    int num=mysql_num_rows(r);
    if(num==0)
    {
        return false;
    }
    MYSQL_ROW row=mysql_fetch_row(r);
    string tk_id=row[0];
    mysql_free_result(r);

    string sql_count =string("select tk_max,tk_count from ticket_info where tk_id=")+tk_id;
    if(mysql_query(&mysql_con,sql_count.c_str())!=0)
    {
        return false;
    }
    r=mysql_store_result(&mysql_con);
    if(r==NULL)
    {
        return false;
    }
    num=mysql_num_rows(r);
    if(num!=1)
    {
        return false;
    }

    row=mysql_fetch_row(r);
    int tk_max=atoi(row[0]);
    int tk_count=atoi(row[1]);
    if(tk_count>0)
    {
        tk_count--;
    }
    mysql_free_result(r);
    Begin();
    //update ticket_info set tk_count=1 where tk_id=1
    string sql_update=string("update ticket_info set tk_count=")+to_string(tk_count)+string(" where tk_id=")+tk_id;
    if(mysql_query(&mysql_con,sql_update.c_str())!=0)
    {
        RollBack();
        return false;
    }
    //delete from ticket_res where res_id=3;
    string sql_del_resid=string("delete from ticket_res where res_id=")+res_id;
    if(mysql_query(&mysql_con,sql_del_resid.c_str())!=0)
    {
        RollBack();
        return false;
    }
    Commit();
    return true;

}
void Accept_Obj::CallBack_Fun()
{
    int c = accept(sockfd, NULL, NULL);
    if (c < 0)
    {
        return;
    }
    Recv_Obj *precv = new Recv_Obj(c, plib);
    if (precv == NULL)
    {
        close(c);
        return;
    }
    plib->MyLibevent_Add(c, precv);
    cout << "accept client:" << c << endl;
}
void Recv_Obj::Send_OK()
{
    Json::Value val;
    val["status"] = "OK";
    send(c, val.toStyledString().c_str(), strlen(val.toStyledString().c_str()), 0);
}
void Recv_Obj::Send_ERR()
{
    Json::Value val;
    val["status"] = "ERR";
    send(c, val.toStyledString().c_str(), strlen(val.toStyledString().c_str()), 0);
}
void Recv_Obj::Send_Json(Json::Value &val)
{
    send(c, val.toStyledString().c_str(), strlen(val.toStyledString().c_str()), 0);
}
void Recv_Obj::User_zc()
{
    string user_tel = m_val["user_tel"].asString();
    string user_name = m_val["user_name"].asString();
    string user_password = m_val["user_password"].asString();

    MysqlClient mysqlcli;
    if (!mysqlcli.Connect_toDb())
    {
        Send_ERR();
        return;
    }
    if (!mysqlcli.Db_user_zc(user_name, user_tel, user_password))
    {
        Send_ERR();
        return;
    }
    Send_OK();
    return;
}
void Recv_Obj::User_dl()
{
    string user_tel = m_val["user_tel"].asString();
    string user_password = m_val["user_password"].asString();

    string user_name;

    MysqlClient mysqlcli;
    if (!mysqlcli.Connect_toDb())
    {
        Send_ERR();
        return;
    }
    if (!mysqlcli.Db_user_dl(user_name, user_tel, user_password))
    {
        Send_ERR();
        return;
    }
    Json::Value val;
    val["status"] = "OK";
    val["user_name"] = user_name;
    Send_Json(val);
}
void Recv_Obj::show_YuYue()
{
    Json::Value resval;
    MysqlClient cli;
    if (!cli.Connect_toDb())
    {
        Send_ERR();
        return;
    }
    if (!cli.Db_show_yuyue(resval))
    {
        Send_ERR();
        return;
    }
    resval["status"] = "OK";
    Send_Json(resval);
    return;
}
void Recv_Obj::User_yd()
{
    string user_tel = m_val["user_tel"].asString();
    string tk_id = m_val["tk_id"].asString();

    MysqlClient cli;
    if (!cli.Connect_toDb())
    {
        Send_ERR();
        return;
    }
    if (!cli.Db_user_yd(user_tel, tk_id))
    {
        Send_ERR();
        return;
    }
    Send_OK();
    return;
}
void Recv_Obj::Show_user_yd()
{
    Json::Value resval;
    string user_tel = m_val["user_tel"].asString();
    MysqlClient cli;
    if (!cli.Connect_toDb())
    {
        Send_ERR();
        return;
    }
    if (!cli.Db_show_yd(resval, user_tel))
    {
        Send_ERR();
        return;
    }
    resval["status"] = "OK";
    Send_Json(resval);
    return;
}
void Recv_Obj::Delete_user_yd()
{
    string res_id = m_val["res_id"].asString();

    MysqlClient cli;
    if (!cli.Connect_toDb())
    {
        Send_ERR();
        return;
    }
    if (!cli.Db_user_delyd(res_id))
    {
        Send_ERR();
        return;
    }
    Send_OK();
    return;
}
void Recv_Obj::CallBack_Fun()
{
    char buff[256] = {0};
    int n = recv(c, buff, 255, 0);
    if (n <= 0)
    {
        plib->MyLibevent_Delete(this);
        delete this;
        return;
    }
    cout << "recv: " << recv << endl;
    m_val.clear();

    Json::Reader Read;
    if (!Read.parse(buff, m_val))
    {
        cout << "json 解析失败" << endl;
        Send_ERR();
        return;
    }
    const int User_Op = m_val["type"].asInt();
    switch (User_Op)
    {
    case DL:
        User_dl();
        break;
    case ZC:
        User_zc();
        break;
    case CKYY:
        show_YuYue();
        break;
    case YD:
        User_yd();
        break;
    case YDXX:
        Show_user_yd();
        break;
    case QXYD:
        Delete_user_yd();
        break;
    }

   /* 
    cout<<"recv: "<<recv<<endl;
    Json::Value val;
    Json::Reader Read;
    Read.parse(buff,val);//反序列化
    cout<<"username:"<<val["user_name"].asString()<<endl;
    cout<<"usertel"<<val["user_tel"].asString()<<endl;
    cout<<"userage"<<val["user_age"].asInt()<<endl;

    val.clear();
    val["status"]="OK";
    val["user_name"]="小白";
    send(c,val.toStyledString().c_str(),strlen(val.toStyledString().c_str()),0);
    */
}
int main()
{
    Socket sock;
    if (!sock.Socket_init())
    {
        cout << "sockfd err" << endl;
        exit(1);
    }
    MyLibevent *plib = new MyLibevent();
    if (plib == NULL) 
    {
        exit(1);
    }
    if (!plib->MyLibevent_Init())
    {
        cout << "MyLibevent init err" << endl;
        exit(1);
    }
    Accept_Obj *pObj = new Accept_Obj(sock.Get_sockfd(), plib);
    if (pObj == NULL)
    {
        exit(1);
    }
    plib->MyLibevent_Add(sock.Get_sockfd(), pObj);
    plib->MyLibevent_Loop();

    delete (pObj);
    delete (plib);
    exit(0);
}