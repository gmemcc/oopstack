import {PageContainer, ProForm, ProFormText} from '@ant-design/pro-components';
import {Card, message} from 'antd';
import React from 'react';
import {postApiHello} from "@/services/oopstack/greeting";

const Welcome: React.FC = () => {
  return (
    <PageContainer>
      <Card>
        <ProForm onFinish={async (values) => {
          let name = values.name;
          let greetResp = await postApiHello({name: name});
          message.success(greetResp.greeting)
        }}>
          <ProFormText name={"name"} label={'Name'}/>
        </ProForm>
      </Card>
    </PageContainer>
  );
};

export default Welcome;
