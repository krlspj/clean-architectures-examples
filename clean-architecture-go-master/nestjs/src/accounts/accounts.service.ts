import { Injectable } from '@nestjs/common';
import { CreateAccountDto } from './dto/create-account.dto';
import { UpdateAccountDto } from './dto/update-account.dto';
import {Account} from "./entities/account.entity";
import {InjectModel} from "@nestjs/sequelize";
import {EmptyResultError, Op} from "sequelize";

@Injectable()
export class AccountsService {
  constructor(
      @InjectModel(Account)
      private accountModel: typeof Account
  ) {}

  create(createAccountDto: CreateAccountDto) {
    return this.accountModel.create(createAccountDto);
  }

  findAll() {
    return this.accountModel.findAll();
  }

  findOne(idOrToken: string) {
    return this.accountModel.findOne({
      where: {
        [Op.or] : {
          id: idOrToken,
          token: idOrToken
        }
      },
      rejectOnEmpty: new EmptyResultError(`Account with ID/TOKEN ${idOrToken} not found`)
    });
  }

  async update(id: string, updateAccountDto: UpdateAccountDto) {
    const account = await this.findOne(id)
    return account.update(updateAccountDto);
  }

  async remove(id: string) {
    const account = await this.findOne(id)
    return account.destroy();
  }
}
