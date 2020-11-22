/* tslint:disable */
/* eslint-disable */
/**
 * GamersNoteAPI
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    Comment,
    CommentFromJSON,
    CommentFromJSONTyped,
    CommentToJSON,
    User,
    UserFromJSON,
    UserFromJSONTyped,
    UserToJSON,
} from './';

/**
 * 
 * @export
 * @interface Article
 */
export interface Article {
    /**
     * 
     * @type {string}
     * @memberof Article
     */
    readonly articleId?: string;
    /**
     * 
     * @type {User}
     * @memberof Article
     */
    author?: User;
    /**
     * 
     * @type {string}
     * @memberof Article
     */
    thumbnailUrl?: string;
    /**
     * 
     * @type {string}
     * @memberof Article
     */
    title?: string;
    /**
     * 
     * @type {string}
     * @memberof Article
     */
    body?: string;
    /**
     * 
     * @type {number}
     * @memberof Article
     */
    readonly likeCount?: number;
    /**
     * 
     * @type {Array<Comment>}
     * @memberof Article
     */
    comments?: Array<Comment>;
    /**
     * 
     * @type {boolean}
     * @memberof Article
     */
    isPublished?: boolean;
    /**
     * 
     * @type {Date}
     * @memberof Article
     */
    readonly publishedAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof Article
     */
    readonly createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof Article
     */
    readonly updatedAt?: Date;
}

export function ArticleFromJSON(json: any): Article {
    return ArticleFromJSONTyped(json, false);
}

export function ArticleFromJSONTyped(json: any, ignoreDiscriminator: boolean): Article {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'articleId': !exists(json, 'article_id') ? undefined : json['article_id'],
        'author': !exists(json, 'author') ? undefined : UserFromJSON(json['author']),
        'thumbnailUrl': !exists(json, 'thumbnail_url') ? undefined : json['thumbnail_url'],
        'title': !exists(json, 'title') ? undefined : json['title'],
        'body': !exists(json, 'body') ? undefined : json['body'],
        'likeCount': !exists(json, 'like_count') ? undefined : json['like_count'],
        'comments': !exists(json, 'comments') ? undefined : ((json['comments'] as Array<any>).map(CommentFromJSON)),
        'isPublished': !exists(json, 'is_published') ? undefined : json['is_published'],
        'publishedAt': !exists(json, 'published_at') ? undefined : (new Date(json['published_at'])),
        'createdAt': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'updatedAt': !exists(json, 'updated_at') ? undefined : (new Date(json['updated_at'])),
    };
}

export function ArticleToJSON(value?: Article | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'author': UserToJSON(value.author),
        'thumbnail_url': value.thumbnailUrl,
        'title': value.title,
        'body': value.body,
        'comments': value.comments === undefined ? undefined : ((value.comments as Array<any>).map(CommentToJSON)),
        'is_published': value.isPublished,
    };
}


