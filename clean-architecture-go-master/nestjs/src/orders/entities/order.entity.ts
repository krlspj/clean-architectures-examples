import {BelongsTo, Column, DataType, ForeignKey, Model, PrimaryKey, Table} from "sequelize-typescript";
import {Account} from "../../accounts/entities/account.entity";

export enum OrderStatus {
    Pending = 'pending',
    Approved = 'approved',
    Rejected = 'rejected'
}

@Table({
    tableName: 'orders',
    createdAt: 'created_at',
    updatedAt: 'created_at',
})
export class Order extends Model{
    @PrimaryKey
    @Column({type: DataType.UUID, defaultValue: DataType.UUIDV4})
    id: string // uuid

    @Column({ allowNull: false, type: DataType.DECIMAL(10, 2) })
    amount: number

    @Column({ allowNull: false })
    credit_card_number: string

    @Column({ allowNull: false })
    credit_card_name: string

    @Column({ allowNull: false, defaultValue: OrderStatus.Pending })
    status: OrderStatus

    @ForeignKey(() => Account)
    @Column({allowNull: false, type: DataType.UUID})
    account_id: string

    @BelongsTo(() => Account)
    account: Account
}
